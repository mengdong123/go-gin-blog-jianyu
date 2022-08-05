package main

import (
	"github.com/mengdong123/go-gin-blog-jianyu/models"
	"github.com/robfig/cron"
	"log"
	"time"
)

// corn 定时任务调用
// 参考资料：https://segmentfault.com/a/1190000023029219
func main() {
	log.Println("Starting...")

	c := cron.New()

	// spec param1：时间格式 param2：调用方法
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
