package main

import (
	"fmt"
	"github.com/sin13cos14/promext-plugin/cron"
)

func main() {
	fmt.Println("promext-plugin start hourly task: insert data to ElasticSearch")
	cron.Start()
}