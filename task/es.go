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
	c, err := elastic.NewClient(elastic.SetURL(config.ElasticURL))
	if err != nil {
		fmt.Errorf("ElaticSearch Client init ERROR: %s", err)
		return
	}
	client = c
}
