package task

import (
	"fmt"
	"github.com/sin13cos14/promext-plugin-es/config"
	"github.com/sin13cos14/promext-plugin-es/data"
	"strings"
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
