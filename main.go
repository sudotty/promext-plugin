package main

import (
	"fmt"
	"github.com/sin13cos14/promext-plugin-es/config"
	"github.com/sin13cos14/promext-plugin-es/handler"
)

func main() {
	handler.Handle()
	fmt.Printf("Finished indexed metric data ,please see it at %s/%s/_search", config.ElasticURL, config.IndexName())
}