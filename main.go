package main

import (
	"fmt"
	"go-xstep/internal/routers"
	"net/http"
	"time"
)

func main() {
	//路由
	r := routers.SetupRouter()
	s := http.Server{
		Addr: ":8000", //端口号
		Handler: r,	  //实现接口handler方法  ServeHTTP(ResponseWriter, *Request)
		ReadTimeout: 30 * time.Second, //请求超时时间
		WriteTimeout: 30 * time.Second, //响应超时时间
		IdleTimeout: 30 * time.Second, //IdleTimeout是启用keep-alives时等待下一个请求的最大时间。如果IdleTimeout为零，则使用ReadTimeout的值。如果两者都是零，则没有超时。
		MaxHeaderBytes: 1 << 20, //header头最大字节数
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
