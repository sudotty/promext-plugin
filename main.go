package main

import (
	"haier.com/promext-plugin-es/handler"
	"fmt"
	"haier.com/promext-plugin-es/config"
)

func main() {
	handler.Handle()
	fmt.Printf("Finished indexed metric data ,please see it at %s/%s/_search", config.ElasticURL, config.IndexName())
}