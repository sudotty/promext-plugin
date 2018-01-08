package config

const (
	CurrentURL      = "/current/metrics?"
	RangeURL        = "/range/metrics?"
	SEPERATOR       = "|"
	STEP            = "1h"
	IndexNamePrefix = "metrics-daily-"
	TypeName        = "metric"
	ElasticURL      = "http://localhost:9200"
	PromexBaseURL   = "http://localhost:8080/api/v1"
)

func IndexName() string {
	return IndexNamePrefix + dayTime()
}
