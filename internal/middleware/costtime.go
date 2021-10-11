package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

//CostTime 花费时间
func CostTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		//当前请求时间
		nowTime := time.Now()
		//处理请求
		c.Next()

		costTime := time.Since(nowTime)
		url := c.Request.URL.String()
		log.Printf("the request URL %s cost %v\n", url, costTime)
	}
}