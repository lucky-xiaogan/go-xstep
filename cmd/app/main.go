package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"go-xstep/config"
	"go-xstep/internal/routers"
	"go-xstep/pkg/cache/xredis"
	"go-xstep/pkg/logger"
	"go-xstep/pkg/x/xreporter"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var (
	envPath  string
	rdb      *redis.Client
	zlogger  *zap.Logger
	reporter *xreporter.Reporter
)

func init() {
	flag.StringVar(&envPath, "env", "./config.yml", "")
}

func main() {
	flag.Parse()
	//config
	env := config.Env(envPath)
	err := config.New(env)
	// 配置文件未正确加载
	if err != nil {
		panic(err)
	}

	//xredis init
	rdb = xredis.NewRedis(config.Conf.Redis.Addr, config.Conf.Redis.Password)

	//初始化 logger
	zlogger, err = logger.NewJSONLogger(
		logger.WithDisableConsole(),
		//logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName, env.Active().Value())),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileP(config.Conf.Logger.File),
	)
	if err != nil {
		panic(err)
	}

	// 用于监听服务退出
	done := make(chan error, 4)
	// 用于控制服务退出，传入同一个 stop，做到只要有一个服务退出了那么另外一个服务也会随之退出
	stop := make(chan struct{})
	// ctx, cannel := context.WithCancel(context.Background())
	// defer cannel()

	go func() {
		done <- AppServer(stop)
	}()

	go func() {
		done <- AdminServer(stop)
	}()

	reporter = xreporter.NewReporter(3, 10)
	go func() {
		reporter.Run(stop)
		done <- nil
	}()

	quit := make(chan os.Signal)
	//设置信号，让程序优雅的退出
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	//TODO 信号处理
	go func() {
		test := make(chan struct{})
		go func() {
			select {
			case <-test:
				fmt.Println("打印不了")
			case <-stop:
				quit <- syscall.SIGINT
			}
		}()
		<-quit
		close(test)
		done <- errors.New("test")
	}()

	var stoped bool
	for i := 0; i < cap(done); i++ {
		err := <-done
		if err != nil {
			fmt.Printf("s.ListenAndServe err: %+v\n", err)
		}

		if !stoped {
			stoped = true
			close(stop)
		}
	}
}

//AppServer 应用
func AppServer(stop <-chan struct{}) error {
	e := routers.NewEntry(config.Conf, rdb, zlogger, reporter)
	r := e.SetupRouter()
	return server(r, config.Conf.Port.HTTPAddr, stop)
}

//AdminServer 管理后台
func AdminServer(stop <-chan struct{}) error {
	//路由
	//e := routers.NewEntry(conf, rdb, zLog)
	//r := e.SetupRouter()
	return server(nil, config.Conf.Port.AdminAddr, stop)
	//go func() {
	//	server(nil, config.Conf.Port.AdminAddr, stop)
	//}()
	//time.Sleep(5 * time.Second)
	//return fmt.Errorf("mock admin exit")
}

func server(handler http.Handler, addr string, stop <-chan struct{}) error {
	s := http.Server{
		Addr:           addr,             //端口号
		Handler:        handler,          //实现接口handler方法  ServeHTTP(ResponseWriter, *Request)
		ReadTimeout:    30 * time.Second, //请求超时时间
		WriteTimeout:   30 * time.Second, //响应超时时间
		IdleTimeout:    30 * time.Second, //IdleTimeout是启用keep-alives时等待下一个请求的最大时间。如果IdleTimeout为零，则使用ReadTimeout的值。如果两者都是零，则没有超时。
		MaxHeaderBytes: 1 << 20,          //header头最大字节数
	}
	go func() {
		<-stop
		// Shutdown 接口，如果没有新的连接了就会释放，传入超时 context
		// 调用这个接口会关闭服务，但是不会中断活动连接
		// 首先会将端口监听移除
		// 然后会关闭所有的空闲连接
		// 然后等待活动的连接变为空闲后关闭
		// 如果等待时间超过了传入的 context 的超时时间，就会强制退出
		// 调用这个接口 server 监听端口会返回 ErrServerClosed 错误
		// 注意，这个接口不会关闭和等待websocket这种被劫持的链接，如果做一些处理。可以使用 RegisterOnShutdown 注册一些清理的方法
		log.Printf("server will exiting, addr: %s", addr)
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}
