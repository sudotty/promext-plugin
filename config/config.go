package config

import (
	"fmt"
	"os"
	"strings"
)

const (
	SEPERATOR       = "|"
	STEP            = "1h"
	UnixStep        = 3600
	IndexNamePrefix = "metrics-hourly-"
	TypeName        = "metric"
	CurrentURL      = "/api/v1/current/metrics?"
	RangeURL        = "/api/v1/range/metrics?"
)

var (
	PromextURL        string
	PromextCurrentURL string
	PromextRangeURL   string
	ESURL             []string
)

func init() {
	PromextURL = os.Getenv("Promext")
	ESURL = strings.Split(os.Getenv("ES"), ",")
	if ESURL == nil {
		fmt.Errorf("ERROR,ESURL is nil,please check your ENV")
	}
	PromextCurrentURL = PromextURL + CurrentURL
	PromextRangeURL = PromextURL + RangeURL
}
func IndexName() string {
	fmt.Println("index time :" + dayTime())
	return IndexNamePrefix + dayTime()
}
