package elasticsearch

import (
	"fmt"
	"log"
	"os"

	"github.com/olivere/elastic"
)

var Client *elastic.Client

func Setup() {
	host := os.Getenv("ELASTICSEARCH_URL")
	port := os.Getenv("ELASTICSEARCH_PORT")

	client, err := elastic.NewClient(
		elastic.SetURL(fmt.Sprintf("http://%s:%s", host, port)),
	)
	if err != nil {
		log.Print(err)
		return
	}

	Client = client
}
