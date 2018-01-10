package cron

import (
	"github.com/robfig/cron"
	"github.com/sin13cos14/promext-plugin/task"
)

func Start() {
	c := cron.New()
	c.AddFunc("@hourly", func() {
		task.Do()
	})
	c.Start()
	select {}
}
