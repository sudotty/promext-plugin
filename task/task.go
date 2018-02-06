package task

import (
	"fmt"

	"github.com/ai-orca/promext-plugin/config"
	"github.com/robfig/cron"
	"os"
	"strconv"
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
	a, _ := strconv.Atoi(os.Getenv("ESStart"))
	b, _ := strconv.Atoi(os.Getenv("ESEnd"))
	if a != 0 && b != 0 {
		config.StartFlag = a
		config.EndFlag = b
		config.IsHistoryJob = true
	} else {
		config.IsHistoryJob = false
	}
	return config.IsHistoryJob
}

func esTask() {
	addMapping()
	bulkIndex()
}
