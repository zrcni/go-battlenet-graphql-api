package battlenet

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string, authToken string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("new request: %v", err)
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken))

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("do request: %v", err)
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("read body: %v", err)
		return nil, err
	}

	return body, nil
}
