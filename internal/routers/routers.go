package routers

import (
	"go-xstep/config"
	"go-xstep/internal/middleware"
	"go-xstep/pkg/x/xreporter"
	"go-xstep/pkg/x/xsort"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type Enter struct {
	conf        *config.Config
	redisClient *redis.Client
	zlog        *zap.Logger
	reporter    *xreporter.Reporter
}

func NewEntry(conf *config.Config, rdb *redis.Client, zlog *zap.Logger, reporter *xreporter.Reporter) *Enter {
	return &Enter{
		conf:        conf,
		redisClient: rdb,
		zlog:        zlog,
		reporter:    reporter,
	}
}

func (e *Enter) SetupRouter() *gin.Engine {
	gin.SetMode("release")
	r := gin.New()
	////gin.default 默认加载日志中间件
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	r.Use(middleware.Logger(e.zlog))
	r.Use(middleware.Recovery(e.zlog, true))
	//自定义中间件
	r.Use(middleware.CostTime())

	//路由
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		//截取/
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})

	//路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hellWorld")
	})

	r.POST("/bubblesort", func(c *gin.Context) {
		var m []int
		c.ShouldBindJSON(&m)
		xsort.SelectedSort(m)
		//fmt.Println(m["n"])
		c.String(http.StatusOK, "%v", m)
	})

	r.GET("/reporter", func(c *gin.Context) {
		for i := 0; i < 100; i++ {
			e.reporter.Report(strconv.Itoa(i))
		}
		c.String(http.StatusOK, "end")
	})

	//默认为监听8080端口
	//r.Run(":8000")
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	return r
}
