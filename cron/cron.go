package cron

import (
	"github.com/robfig/cron"
	"github.com/sin13cos14/promext-plugin-es/handler"
)

func Task() {
	c := cron.New()
	c.AddFunc("@hourly", func() {
		handler.Handle()
	})
	c.Start()
	select {}
}
