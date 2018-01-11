package cron

import (
	"github.com/ai-orca/promext-plugin/task"
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
