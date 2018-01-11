package task

import (
	"fmt"
	"strings"

	"git.haier.net/monitor/promext-apm-plugin/config"
	"git.haier.net/monitor/promext-apm-plugin/data"
)

func Do() {
	for k, v := range data.ProcessMetricData() {
		kArr := strings.Split(k, config.SEPERATOR)
		m := getMetricES(kArr[1], kArr[0], v)
		indexMetricMapping()
		indexMetricToES(m)
	}
	fmt.Printf("Finished indexed metric data ,please see it at %s/%s/_search", config.ElasticURL, config.IndexName)
}
