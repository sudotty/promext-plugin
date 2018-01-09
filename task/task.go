package task

import (
	"fmt"
	"github.com/sin13cos14/promext-plugin-es/config"
	"github.com/sin13cos14/promext-plugin-es/storage"
	"strings"
)

func Do() {
	for key, value := range storage.ProcessMetricData() {
		keyArray := strings.Split(key, config.SEPERATOR)
		metricToES := getMetricES(keyArray[1], keyArray[0], value)
		indexMetricMapping()
		indexMetricToES(metricToES)
	}
	fmt.Printf("Finished indexed metric data ,please see it at %s/%s/_search", config.ElasticURL, config.IndexName())
}
