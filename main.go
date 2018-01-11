package main

import (
	"fmt"

	"promext-apm-plugin/cron"
)

func main() {
	fmt.Println("promext-plugin start hourly task: insert data to ElasticSearch")
	cron.Start()
}
