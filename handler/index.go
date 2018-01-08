package handler

import (
	"context"
	"fmt"
	"github.com/sin13cos14/promext-plugin-es/config"
	"github.com/sin13cos14/promext-plugin-es/storage"
	"strings"
	"time"
)

type MetricES struct {
	Ip      string            `json:"ip"`
	Project string            `json:"project"`
	Values  map[string]string `json:"values"`
}

var ctx = context.Background()

func indexMetricToES(metricES *MetricES) {

	_, err := client().Index().
		Index(config.IndexName()).
		Type(config.TypeName).
		BodyJson(metricES).
		Timestamp(time.Now().Format("200601021504")).
		Do(ctx)
	if err != nil {
		fmt.Errorf("Index Metric To ES  ERROR: %s/n", err)
		return
	}
}

func getMetricES(ip, project string, value map[string]string) *MetricES {
	return &MetricES{
		Ip:      ip,
		Project: project,
		Values:  value,
	}
}

func Handle() {
	for key, value := range storage.ProcessMetricData() {
		keyArray := strings.Split(key, config.SEPERATOR)
		metricToES := getMetricES(keyArray[1], keyArray[0], value)
		indexMetricMapping()
		indexMetricToES(metricToES)
	}
	fmt.Printf("Finished indexed metric data ,please see it at %s/%s/_search", config.ElasticURL, config.IndexName())
}
