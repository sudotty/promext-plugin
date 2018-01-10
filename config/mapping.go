package config

const Mapping = `{
    "settings": {
        "number_of_shards": 6,
        "number_of_replicas": 1
    },
    "mappings": {
        "metric": {
            "_all": {
                "enabled": false
            },
            "properties": {
                "ip": {
                    "type": "ip"
                },
                "project": {
                    "type": "keyword"
                },
                "values": {
                    "properties": {
                        "cpuUtilization": {
                            "type": "double"
                        },
                        "cpuUtilizationAvg": {
                            "type": "double"
                        },
						"cpuUtilizationMedian": {
                            "type": "double"
                        },
                        "cpuUtilizationMax": {
                            "type": "double"
                        },
                        "cpuUtilizationMin": {
                            "type": "double"
                        },
                        "diskUtilization": {
                            "type": "double"
                        },
                        "diskUtilizationMax": {
                            "type": "double"
                        },
                        "memoryUtilization": {
                            "type": "double"
                        },
                        "memoryUtilizationAvg": {
                            "type": "double"
                        },
						"memoryUtilizationMedian": {
                            "type": "double"
                        },
                        "memoryUtilizationMax": {
                            "type": "double"
                        },
                        "memoryUtilizationMin": {
                            "type": "double"
                        }
                    }
                }
            }
        }
    }
}`
