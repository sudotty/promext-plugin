package task

import (
	"github.com/ai-orca/promext-plugin/config"
	"github.com/robfig/cron"
)

func Cron() {
	c := cron.New()
	c.AddFunc("@hourly", func() {
		RunBulkTask()
	})
	c.Start()
	select {}
}

func HistoryData() {
	t := config.Times()
	for i := 0; i < t; i++ {
		RunBulkTask()
		config.Increase()
	}
}
