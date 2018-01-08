package handler

import (
	"fmt"
	"github.com/sin13cos14/promext-plugin-es/config"
	"gopkg.in/olivere/elastic.v5"
)

func client() *elastic.Client {
	client, err := elastic.NewClient(elastic.SetURL(config.ElasticURL))
	if err != nil {
		fmt.Errorf("ElaticSearch Client init ERROR: %s", err)
		return nil
	}
	return client
}
