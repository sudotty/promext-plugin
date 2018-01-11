package main

import (
	"fmt"

	"github.com/ai-orca/promext-plugin/cron"
)

func main() {
	fmt.Println("promext-plugin start hourly task: insert data to ElasticSearch")
	cron.Start()
}
