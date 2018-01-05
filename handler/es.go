package handler

import (
	"gopkg.in/olivere/elastic.v5"
	"fmt"
	"haier.com/promext-plugin-es/config"
)

func client() *elastic.Client {
	client, err := elastic.NewClient(elastic.SetURL(config.ElasticURL))
	if err != nil {
		fmt.Println(err)
	}
	return client
}