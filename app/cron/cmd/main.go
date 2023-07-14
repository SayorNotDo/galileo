package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New() // 新建定时任务对象

	_, err := c.AddFunc("@every 1s", func() {
		fmt.Println("tick every 1 second")
	}) // 给对象增加定时任务
	if err != nil {
		return
	}

	c.Start()
	time.Sleep(time.Second * 5)
}
