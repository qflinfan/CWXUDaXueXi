package main

import (
	"Daxuexi/Request"
	"github.com/go-co-op/gocron"
	"time"
)

func main() {
	s := gocron.NewScheduler(time.UTC)
	// 每周一早上9点执行函数
	Request.Task()
	s.Every(1).Monday().At("09:00").Do(Request.Task)
	s.StartBlocking()
}
