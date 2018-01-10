package task

import "github.com/sin13cos14/promext-plugin/config"

func indexMetricMapping() {
	exists, err := client.IndexExists(config.IndexName).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err := client.CreateIndex(config.IndexName).BodyString(config.Mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
	}
}
