package data

import (
	"encoding/json"
	"fmt"

	"git.haier.net/monitor/promext-apm-plugin/config"
)

type MetricNestedMap map[string]map[string]string

func metricNestedMapKey(m *Metric) string {
	return m.Project + config.SEPERATOR + m.Host
}
func (mmap MetricNestedMap) currentDataTransform(data []MetricModelCurrent) MetricNestedMap {
	for _, value := range data {
		m := &value.Metric
		k := metricNestedMapKey(m)
		if mmap[k] == nil {
			mmap[k] = make(map[string]string)
		}
		v := value.Value[1]
		if v == "NaN" {
			continue
		}
		mmap[k][m.Name] = v.(string)
	}
	return mmap
}
func (mmap MetricNestedMap) rangeDataTransform(data []MetricModelRange) MetricNestedMap {
	for _, value := range data {
		m := &value.Metric
		k := metricNestedMapKey(m)
		if mmap[k] == nil {
			mmap[k] = make(map[string]string)
		}
		v := value.Value[0].([]interface{})[1]
		if v == "NaN" {
			continue
		}
		mmap[k][value.Metric.Name] = v.(string)
	}
	return mmap
}

func ProcessMetricData() MetricNestedMap {
	mmap := MetricNestedMap{}
	var mc MetricsCurrent
	var mr MetricsRange
	currentdata := promextData(currentdataURL())
	if currentdata == nil {
		fmt.Errorf("CurrentData is nil ERROR")
		return mmap
	}
	err := json.Unmarshal(currentdata, &mc)
	if err != nil {
		fmt.Errorf("CurrentData Unmarshal ERROR: %s", err)
		return mmap
	}
	rangeData := promextData(rangeDataURL())
	if rangeData == nil {
		fmt.Errorf("rangeData is nil ERROR")
		return mmap
	}
	err2 := json.Unmarshal(rangeData, &mr)
	if err2 != nil {
		fmt.Errorf("RangeData Unmarshal ERROR: %s", err)
		return mmap
	}
	return mmap.currentDataTransform(mc.Data).rangeDataTransform(mr.Data)
}
