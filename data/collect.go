package data

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"git.haier.net/monitor/promext-apm-plugin/config"
)

func promextParam(start, end, step string) string {
	return "start=" + start + "&end=" + end + "&step=" + step
}

func currentdataURL() string {
	start, end, _ := config.GetStartEndTime()
	return config.PromextCurrentURL + promextParam(start, end, config.STEP)

}

func rangeDataURL() string {
	start, end, _ := config.GetStartEndTime()
	return config.PromextRangeURL + promextParam(start, end, config.STEP)
}

type Metric struct {
	Name             string `json:"_name_"`
	App              string `json:"app"`
	Cluster          string `json:"cluster"`
	ExportedInstance string `json:"exported_instance"`
	Host             string `json:"host"`
	Instance         string `json:"instance"`
	Job              string `json:"job"`
	Project          string `json:"project"`
	Url              string `json:"url"`
	Zone             string `json:"zone"`
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
	r, err := http.Get(url)
	if r == nil {
		fmt.Errorf("get resp data ERROR : %s", err)
		return nil
	}
	if err != nil {
		fmt.Errorf("get API data ERROR : %s", err)
		return nil
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Errorf("IO read data ERROR: %s", err)
		return nil
	}
	return body
}
