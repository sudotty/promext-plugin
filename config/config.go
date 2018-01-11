package config

const (
	SEPERATOR       = "|"
	STEP            = "1h"
	IndexNamePrefix = "metrics-daily-"
	TypeName        = "metric"
	CurrentURL      = "/api/v1/current/metrics?"
	RangeURL        = "/api/v1/range/metrics?"
	PromextBaseURL  = "http://localhost:8080"
	ElasticURL      = "http://localhost:9200"
)

var (
	IndexName         string
	PromextCurrentURL string
	PromextRangeURL   string
)

func init() {
	IndexName = IndexNamePrefix + dayTime()
	PromextCurrentURL = PromextBaseURL + CurrentURL
	PromextRangeURL = PromextBaseURL + RangeURL
}
