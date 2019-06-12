package main

import (
	"log"
	"fmt"
	"os"
	"time"
	"gopkg.in/olivere/elastic.v6"
)

var host = "http://172.20.21.113:9200/"

func init() {
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

	fmt.Println("connect es success!", client)
}

func main() {

}