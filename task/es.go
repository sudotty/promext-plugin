package task

import (
	"context"
	"fmt"

	"github.com/ai-orca/promext-plugin/config"
	"gopkg.in/olivere/elastic.v5"
)

var client *elastic.Client
var ctx = context.Background()

func init() {
	c, err := elastic.NewClient(setURL())
	if err != nil {
		fmt.Errorf("ElaticSearch Client init ERROR: %s", err)
		return
	}
	client = c
}
func setURL() elastic.ClientOptionFunc {
	if config.ElasticURLNode1 != "" && config.ElasticURLNode2 == "" && config.ElasticURLNode3 == "" {
		return elastic.SetURL(config.ElasticURLNode1)
	} else if config.ElasticURLNode1 != "" && config.ElasticURLNode2 != "" && config.ElasticURLNode3 == "" {
		return elastic.SetURL(config.ElasticURLNode1, config.ElasticURLNode2)
	} else if config.ElasticURLNode1 != "" && config.ElasticURLNode2 != "" && config.ElasticURLNode3 != "" {
		return elastic.SetURL(config.ElasticURLNode1, config.ElasticURLNode2, config.ElasticURLNode3)
	} else if config.ElasticURLNode1 != "" && config.ElasticURLNode2 == "" && config.ElasticURLNode3 != "" {
		return elastic.SetURL(config.ElasticURLNode1, config.ElasticURLNode3)
	} else {
		return nil
	}
}
