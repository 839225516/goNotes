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
	fmt.Println("client: ", client)

	// 删除es索引
	ctx := context.Background()

	exist, err := client.IndexExists("twitter").Do(ctx)

	if err != nil {
		// pass
	}

	if !exist {
		// 不存在，创建索引
	} else {
		// 删除索引
		deleteIndex, err := client.DeleteIndex("twitter").Do(ctx)
		if err != nil {
			panic(err)
		}

		if !deleteIndex.Acknowledged {
			fmt.Println("!deleteIndex.Acknowledged")
		} else {
			fmt.Println("deleteIndex.Acknowledged")
		}

	}

}
