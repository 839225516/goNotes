package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/olivere/elastic.v6"
)

var host = "http://172.20.21.113:9200/"

func connect() *elastic.Client {
	client, err := elastic.NewClient(
		elastic.SetURL(host),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetGzip(true),
		elastic.SetErrorLog(log.New(os.Stderr, "Elastic ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)

	if err != nil {
		panic(err)
	}

	return client
}

func main() {
	client := connect()
	fmt.Println("client:", client)

	mapping := `{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"tweet":{
			"properties":{
				"tags":{
					"type":"text"
				},
				"location":{
					"type":"geo_point"
				}
			}
		}
	}
}`

	ctx := context.Background()

	createIndex, err := client.CreateIndex("twitter").BodyString(mapping).Do(ctx)

	if err != nil {
		panic(err)
	}

	if !createIndex.Acknowledged {
		fmt.Println("!createIndex.Acknowledged")
	} else {
		fmt.Println("createIndex.Acknowledged")
	}

}
