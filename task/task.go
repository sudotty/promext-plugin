package task

import (
	"flag"
	"fmt"

	"github.com/ai-orca/promext-plugin/config"
	"github.com/robfig/cron"
)

func init() {
	initMode()
	initES()
}

func Run() {
	if config.IsHistoryJob {
		fmt.Println("history task: insert data to elastic")
		cycleTask()
	} else {
		fmt.Println("hourly task: insert data to elastic")
		cronTask()
	}
}

func cronTask() {
	c := cron.New()
	c.AddFunc("@hourly", func() {
		esTask()
	})
	c.Start()
	select {}
}

func cycleTask() {
	t := config.Cycles()
	for i := 0; i < t; i++ {
		esTask()
		config.Increase()
	}
}

func initMode() bool {
	flag.IntVar(&config.StartFlag, "start", 0, "start time")
	flag.IntVar(&config.EndFlag, "end", 0, "end time")
	flag.Parse()
	config.IsHistoryJob = config.StartFlag != 0 && config.EndFlag != 0
	return config.IsHistoryJob
}

func esTask() {
	addMapping()
	bulkIndex()
}
func EsTask() {
	addMapping()
	bulkIndex()
}
