package task

import (
	"fmt"
	"github.com/sin13cos14/promext-plugin-es/config"
	"time"
)

type MetricES struct {
	Ip      string            `json:"ip"`
	Project string            `json:"project"`
	Values  map[string]string `json:"values"`
}

func indexMetricToES(metricES *MetricES) {
	_, err := client.Index().
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
