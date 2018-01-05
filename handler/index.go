package handler

import (
	"context"
	"time"
	"fmt"
	"strings"
	"haier.com/promext-plugin-es/config"
	"haier.com/promext-plugin-es/storage"
)

type MetricES struct {
	Ip      string            `json:"ip"`
	Time    string            `json:"time"`
	Project string            `json:"project"`
	Values  map[string]string `json:"values"`
}

func indexMetricToES(metricES *MetricES) {
	_, err := client().Index().
		Index(config.IndexName()).
		Type(config.TypeName).
		BodyJson(metricES).
		Timestamp(time.Now().Format(time.RFC3339)).
		Do(context.Background())
	if err != nil {
		fmt.Errorf("Index Metric To ES  ERROR: %s/n", err)
		return
	}
}

func getMetricES(ip, project string, value map[string]string) *MetricES {
	return &MetricES{
		Ip:      ip,
		Time:    config.MetricTime(),
		Project: project,
		Values:  value,
	}
}

func Handle() {
	for key, value := range storage.ProcessMetricData() {
		keyArray := strings.Split(key, config.SEPERATOR)
		metricToES := getMetricES(keyArray[1], keyArray[0], value)
		indexMetricToES(metricToES)
	}
}
