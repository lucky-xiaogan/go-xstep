package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	//gin.default 默认加载日志中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//路由
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		//截取/
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})
	//默认为监听8080端口
	//r.Run(":8000")
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	return r
}
