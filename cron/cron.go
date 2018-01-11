package cron

import (
	"git.haier.net/monitor/promext-apm-plugin/task"
	"github.com/robfig/cron"
)

func Start() {
	c := cron.New()
	c.AddFunc("@hourly", func() {
		task.Do()
	})
	c.Start()
	select {}
}
