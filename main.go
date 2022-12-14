package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/mengdong123/go-gin-blog-jianyu/models"
	"github.com/mengdong123/go-gin-blog-jianyu/pkg/logging"
	"github.com/mengdong123/go-gin-blog-jianyu/pkg/setting"
	"github.com/mengdong123/go-gin-blog-jianyu/routers"
	"log"
	"syscall"
)

func main() {

	// 将init方法更改为Setup方法，避免多 init 的情况，尽量由程序把控初始化的先后顺序
	setting.Setup()
	models.Setup()
	logging.Setup()

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
