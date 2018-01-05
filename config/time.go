package config

import (
	"strconv"
	"time"
)

func GetStartEndTime() (string, string, string) {
	now := time.Now()
	nowRight := now.Format("2006-01-02 15:00")
	metricTime, _ := time.ParseInLocation("2006-01-02 15:00", nowRight, time.Local)
	start := metricTime.Unix()
	end := metricTime.Unix() + 3600
	return strconv.FormatInt(start, 10), strconv.FormatInt(end, 10), metricTime.Format("200601021504")
}

func dayTime() string {
	return time.Now().Format("2006-01-02")
}
func MetricTime() string {
	return time.Now().Format("200601021500")
}
