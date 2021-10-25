package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-xstep/config"
	"go-xstep/internal/routers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var envPath string
var conf *config.Config

func init() {
	flag.StringVar(&envPath, "env", "./config.yml", "")
}

func main() {
	flag.Parse()
	env := config.Env(envPath)
	conf = config.New(env)

	done := make(chan error, 2)
	stop := make(chan struct{})
	go func() {
		done <- httpServer(stop)
	}()

	go func() {
		done <- AdminServer(stop)
	}()

	quit := make(chan os.Signal)
	//设置信号，让程序优雅的退出
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	var stoped bool
	for i := 0; i < cap(done); i++ {
		select {
		case <-quit:
			fmt.Println("signal notify shutdown")
		case err := <-done:
			if err != nil {
				fmt.Printf("s.ListenAndServe err: %v", err)
			}
		}

		if !stoped {
			stoped = true
			close(stop)
		}
	}
}

func httpServer(stop <-chan struct{}) error {
	//路由
	gin.SetMode(gin.DebugMode)
	r := routers.SetupRouter()
	s := http.Server{
		Addr:           conf.Port.HTTPAddr, //端口号
		Handler:        r,                  //实现接口handler方法  ServeHTTP(ResponseWriter, *Request)
		ReadTimeout:    30 * time.Second,   //请求超时时间
		WriteTimeout:   30 * time.Second,   //响应超时时间
		IdleTimeout:    30 * time.Second,   //IdleTimeout是启用keep-alives时等待下一个请求的最大时间。如果IdleTimeout为零，则使用ReadTimeout的值。如果两者都是零，则没有超时。
		MaxHeaderBytes: 1 << 20,            //header头最大字节数
	}
	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}

func AdminServer(stop <-chan struct{}) error {
	//路由
	gin.SetMode(gin.DebugMode)
	r := routers.SetupRouter()
	s := http.Server{
		Addr:           conf.Port.AdminAddr, //端口号
		Handler:        r,                   //实现接口handler方法  ServeHTTP(ResponseWriter, *Request)
		ReadTimeout:    30 * time.Second,    //请求超时时间
		WriteTimeout:   30 * time.Second,    //响应超时时间
		IdleTimeout:    30 * time.Second,    //IdleTimeout是启用keep-alives时等待下一个请求的最大时间。如果IdleTimeout为零，则使用ReadTimeout的值。如果两者都是零，则没有超时。
		MaxHeaderBytes: 1 << 20,             //header头最大字节数
	}
	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}
