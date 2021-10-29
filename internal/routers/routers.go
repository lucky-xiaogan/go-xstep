package routers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-xstep/internal/middleware"
	"go-xstep/pkg/cache/redis"
	"go-xstep/pkg/x/xsort"
	"log"
	"net/http"
	"strings"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	//gin.default 默认加载日志中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
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

	r.POST("/bubblesort", func(c *gin.Context) {
		var m []int
		c.ShouldBindJSON(&m)
		xsort.SelectedSort(m)
		//fmt.Println(m["n"])
		c.String(http.StatusOK, "%v", m)
	})

	r.GET("/redis", func(c *gin.Context) {

		val, err := redis.RedisDB.Get(context.Background(), "test").Result()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(val)


	})

	r.GET("/", func(c *gin.Context) {

		test, err := redis.RedisDB.Get(context.Background(), "test").Result()
		if err != nil {
			fmt.Println("hello")
			log.Fatal(err)
		}
		fmt.Println(test)
		val := redis.RedisDB.Set(context.Background(), "test", 1, 60)
		fmt.Println(val)

	})

	//默认为监听8080端口
	//r.Run(":8000")
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	return r
}
