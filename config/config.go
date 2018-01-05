package config

const (
	PromexBaseURL   = "http://localhost:8080/api/v1"
	CurrentURL      = "/current/metrics?"
	RangeURL        = "/range/metrics?"
	SEPERATOR       = "|"
	STEP            = "1h"
	IndexNamePrefix = "metrics-daily-"
	TypeName        = "metric"
	ElasticURL      = "http://localhost:9200"
)

func IndexName() string {
	return IndexNamePrefix + dayTime()
}
