package task

import "git.haier.net/monitor/promext-apm-plugin/config"

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
