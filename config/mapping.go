package config

const Mapping = `{
    "settings": {
        "number_of_shards": 6,
        "number_of_replicas": 1
    },
    "mappings": {
        "_default_": {
            "dynamic_templates": [{
                "values": {
                    "path_match": "values.*",
                    "mapping": {
                        "type": "double"
                    }
                }
            }, {
                "time": {
                    "match": "*time",
                    "mapping": {
                    "type": "date",
                    "format": "epoch_second"
                    }
                }
            }, {
               "default":{
                    "match": "*",
                    "match_mapping_type": "string",
                    "mapping": {
                        "type": "string",
                        "index": "not_analyzed"
                    }
                }
            }],
            "properties": {
            
            }
        }      
    }
}`
