package storage

import (
	"encoding/json"
	"fmt"
	"github.com/sin13cos14/promext-plugin-es/config"
	"io/ioutil"
	"net/http"
)

func promextParam(start, end, step string) string {
	return "start=" + start + "&end=" + end + "&step=" + step
}

func currentdataURL() string {
	start, end, _ := config.GetStartEndTime()
	return config.PromexBaseURL + config.CurrentURL + promextParam(start, end, config.STEP)

}

func rangeDataURL() string {
	start, end, _ := config.GetStartEndTime()
	return config.PromexBaseURL + config.RangeURL + promextParam(start, end, config.STEP)

}

type Metric struct {
	Name     string `json:"_name_"`
	Cluster  string `json:"cluster"`
	Instance string `json:"instance"`
	Job      string `json:"job"`
	Project  string `json:"project"`
}
type Value []interface{}

type MetricModelCurrent struct {
	Metric Metric `json:"metric"`
	Value  Value  `json:"value"`
}
type MetricModelRange struct {
	Metric Metric `json:"metric"`
	Value  Value  `json:"values"`
}
type MetricsCurrent struct {
	Status string               `json:"status"`
	Data   []MetricModelCurrent `json:"data"`
}
type MetricsRange struct {
	Status string             `json:"status"`
	Data   []MetricModelRange `json:"data"`
}

func promextData(url string) []byte {
	resp, err := http.Get(url)
	if resp == nil {
		fmt.Errorf("get resp data ERROR : %s", err)
		return nil
	}
	if err != nil {
		fmt.Errorf("get API data ERROR : %s", err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Errorf("IO read data ERROR: %s", err)
		return nil
	}
	return body
}

type MetricNestedMap map[string]map[string]string

func currentDataTransform(metricsMap MetricNestedMap, data []MetricModelCurrent) MetricNestedMap {
	for _, value := range data {
		key := value.Metric.Project + config.SEPERATOR + value.Metric.Instance
		if metricsMap[key] == nil {
			metricsMap[key] = make(map[string]string)
		}
		v := value.Value[1]
		if v == "NaN" {
			continue
		}
		metricsMap[key][value.Metric.Name] = v.(string)
	}
	return metricsMap
}

func rangeDataTransform(metricsMap MetricNestedMap, data []MetricModelRange) MetricNestedMap {
	for _, value := range data {
		key := value.Metric.Project + config.SEPERATOR + value.Metric.Instance
		if metricsMap[key] == nil {
			metricsMap[key] = make(map[string]string)
		}
		v := value.Value[0].([]interface{})[1]
		if v == "NaN" {
			continue
		}
		metricsMap[key][value.Metric.Name] = v.(string)
	}
	return metricsMap
}

func ProcessMetricData() MetricNestedMap {
	metricsMap := MetricNestedMap{}
	var metricsCurrent MetricsCurrent
	var metricsRange MetricsRange
	currentdata := promextData(currentdataURL())
	if currentdata == nil {
		fmt.Errorf("CurrentData is nil ERROR")
		return metricsMap
	}
	err := json.Unmarshal(currentdata, &metricsCurrent)
	if err != nil {
		fmt.Errorf("CurrentData Unmarshal ERROR: %s", err)
		return metricsMap
	}
	rangeData := promextData(rangeDataURL())
	if rangeData == nil {
		fmt.Errorf("rangeData is nil ERROR")
		return metricsMap
	}
	err2 := json.Unmarshal(rangeData, &metricsRange)
	if err2 != nil {
		fmt.Errorf("RangeData Unmarshal ERROR: %s", err)
		return metricsMap
	}
	metricsMap = currentDataTransform(metricsMap, metricsCurrent.Data)
	metricsMap = rangeDataTransform(metricsMap, metricsRange.Data)
	return metricsMap
}
