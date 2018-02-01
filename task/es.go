package task

import (
	"context"
	"fmt"

	"github.com/ai-orca/promext-plugin/config"
	"github.com/ai-orca/promext-plugin/data"
	"gopkg.in/olivere/elastic.v5"
)

var client *elastic.Client
var ctx = context.Background()

func initES() {
	c, err := elastic.NewClient(elastic.SetURL(config.ESURL...))
	if err != nil {
		fmt.Errorf("ElaticSearch Client init ERROR: %s", err)
		return
	}
	client = c
}
func addMapping() {
	exists, err := client.IndexExists(config.IndexName()).Do(ctx)
	if err != nil {
		fmt.Println("ERROR,index exist ERROR: ", err)
		return
	}
	if !exists {
		_, err := client.CreateIndex(config.IndexName()).BodyString(config.Mapping).Do(ctx)
		if err != nil {
			fmt.Println("ERROR,add mapping ERROR: ", err)
			return
		}
	}
}
func bulkIndex() {
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
		fmt.Errorf("Error %s indexed metric data , %s/%s/_search\n", err, config.ESURL[0], config.IndexName())
	}
	fmt.Printf("Finished indexed metric data ,please see it at %s/%s/_search\n", config.ESURL[0], config.IndexName())
}
