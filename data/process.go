package data

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ai-orca/promext-plugin/config"
	"os"
)

type MetricES struct {
	Project string            `json:"project"`
	Ip      string            `json:"ip"`
	Zone    string            `json:"zone"`
	Ctime   string            `json:"ctime"`
	Values  map[string]string `json:"values"`
}
type MetricNestedMap map[string]*MetricES

func GetKey(m *Metric) string {
	return m.Project + config.SEPERATOR + m.Instance
}

func getValue(v Value, t string) (float64, string) {
	if t == "current" {
		return v[0].(float64), v[1].(string)
	} else {
		return v[0].([]interface{})[0].(float64), v[0].([]interface{})[1].(string)
	}
}

func metricES(m *Metric, v string, ct float64) *MetricES {
	metricValues := make(map[string]string)
	metricValues[m.Name] = v
	return &MetricES{
		Project: m.Project,
		Ip:      m.App,
		Ctime:   strconv.FormatFloat(ct, 'f', 0, 64),
		Values:  metricValues,
		Zone:    os.Getenv("Zone"),
	}
}
func (mp MetricNestedMap) currentDataTransform(data []MetricModelCurrent) MetricNestedMap {
	for _, value := range data {
		ct, v := getValue(value.Value, "current")
		if v == "NaN" {
			continue
		}
		m := &value.Metric
		k := GetKey(m)
		if mp[k] == nil {
			mp[k] = metricES(m, v, ct)
		} else {
			mp[k].Values[m.Name] = v
		}
	}
	return mp
}

func (mp MetricNestedMap) rangeDataTransform(data []MetricModelRange) MetricNestedMap {
	for _, value := range data {
		ct, v := getValue(value.Value, "range")
		if v == "NaN" {
			continue
		}
		m := &value.Metric
		k := GetKey(m)
		if mp[k] == nil {
			mp[k] = metricES(m, v, ct)
		} else {
			mp[k].Values[m.Name] = v
		}
	}
	return mp
}

func ProcessMetricData() MetricNestedMap {
	var mc MetricsCurrent
	var mr MetricsRange
	mp := MetricNestedMap{}
	currentdata := promextData(getURL(config.PromextCurrentURL))
	if currentdata == nil {
		fmt.Errorf("CurrentData is nil ERROR")
		return mp
	}
	err := json.Unmarshal(currentdata, &mc)
	if err != nil {
		fmt.Errorf("CurrentData Unmarshal ERROR: %s", err)
		return mp
	}
	rangeData := promextData(getURL(config.PromextRangeURL))
	if rangeData == nil {
		fmt.Errorf("rangeData is nil ERROR")
		return mp
	}
	err2 := json.Unmarshal(rangeData, &mr)
	if err2 != nil {
		fmt.Errorf("RangeData Unmarshal ERROR: %s", err)
		return mp
	}
	return mp.currentDataTransform(mc.Data).rangeDataTransform(mr.Data)
}
