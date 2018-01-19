package task

import (
	"fmt"

	"github.com/ai-orca/promext-plugin/config"
	"github.com/ai-orca/promext-plugin/data"
	"gopkg.in/olivere/elastic.v5"
)

func runBulkTask() {
	indexMetricMapping()
	bulkRequest := client.Bulk()
	for _, v := range data.ProcessMetricData() {
		indexReq := elastic.NewBulkIndexRequest().
			Index(config.IndexName()).
			Type(config.TypeName).
			Id(v.Project + config.SEPERATOR + v.Ip + config.SEPERATOR + v.Ctime).
			Doc(v)
		bulkRequest.Add(indexReq)
	}
	_, err := bulkRequest.Do(ctx)
	if err != nil {
		fmt.Errorf("Error %s indexed metric data , %s/%s/_search\n", err, config.ElasticURL, config.IndexName())
	}
	fmt.Printf("Finished indexed metric data ,please see it at %s/%s/_search\n", config.ElasticURL, config.IndexName())
}
