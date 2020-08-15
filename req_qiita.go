package main

import (
	"fmt"
	"io/ioutil"

	// "io/ioutil"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	token := os.Getenv("Q_TOKEN")
	log.Print(token)
	reqURL := "https://qiita.com/api/v2/authenticated_user/items"
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		log.Print("error")
		return
	}
	req.Header.Add(
		"Authorization", fmt.Sprintf("Bearer %s", token),
	)
	values := url.Values{}
	values.Add("page", "1")
	values.Add("per_page", "20")
	req.URL.RawQuery = values.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print("req error")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	type resData struct {
		titles string
	}
	var data interface{}
	json.Unmarshal(body, &data)

	if err == nil {
		log.Print(data)
	}
}
