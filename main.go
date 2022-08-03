package main

import (
	"fmt"
	"github.com/mengdong123/go-gin-blog-jianyu/pkg/setting"
	"github.com/mengdong123/go-gin-blog-jianyu/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
