package main

import (
	"flag"
	"fmt"

	"github.com/ai-orca/promext-plugin/config"
	"github.com/ai-orca/promext-plugin/task"
)

func checkTaskType() bool {
	flag.IntVar(&config.StartFlag, "start", 0, "start time")
	flag.IntVar(&config.EndFlag, "end", 0, "end time")
	flag.Parse()
	config.IsHistoryJob = config.StartFlag != 0 && config.EndFlag != 0
	return config.IsHistoryJob
}

func main() {
	checkTaskType()
	if config.IsHistoryJob {
		fmt.Println("history task: insert data to elastic")
		task.HistoryData()
	} else {
		fmt.Println("hourly task: insert data to elastic")
		task.Cron()
	}
}
