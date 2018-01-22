package config

import "os"

const (
	SEPERATOR       = "|"
	STEP            = "1h"
	IndexNamePrefix = "metrics-daily-"
	TypeName        = "metric"
	CurrentURL      = "/api/v1/current/metrics?"
	RangeURL        = "/api/v1/range/metrics?"
)

var (
	PromextURL        string
	ElasticURLNode1   string
	ElasticURLNode2   string
	ElasticURLNode3   string
	PromextCurrentURL string
	PromextRangeURL   string
)

func init() {
	PromextURL = os.Getenv("PromextURL")
	ElasticURLNode1 = os.Getenv("ESURL1")
	ElasticURLNode2 = os.Getenv("ESURL2")
	ElasticURLNode3 = os.Getenv("ESURL3")
	PromextCurrentURL = PromextURL + CurrentURL
	PromextRangeURL = PromextURL + RangeURL
}
func IndexName() string {
	return IndexNamePrefix + dayTime()
}
