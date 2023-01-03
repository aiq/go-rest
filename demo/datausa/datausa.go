package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aiq/go-rest"
)

type populationData struct {
	Nation     string `json:"Nation"`
	Year       string `json:"Year"`
	Population int64  `json:"Population"`
}

type resp struct {
	Data []populationData `json:"data"`
}

// example from https://datausa.io/about/api/
func main() {
	c := rest.NewJSONClient(rest.Wrap(func(req *rest.Request) ([]byte, error) {
		url, err := rest.URL("https://datausa.io/api", req.Path, req.Parameters)
		if err != nil {
			return []byte{}, err
		}
		log.Printf("%s from %v", req.Method, url)
		client := &http.Client{}
		resp, err := client.Do(&http.Request{
			Method: req.Method,
			URL:    url,
			Body:   rest.Body(req.Body),
		})
		if err != nil {
			return []byte{}, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return []byte{}, fmt.Errorf("unexpected HTTP Status Code %d from %q", resp.StatusCode, url.String())
		}

		return rest.ReadBytes(resp.Body)
	}))

	var r resp
	err := c.GetJSON("data", &r, "drilldowns", "Nation", "measures", "Population")
	if err != nil {
		log.Fatalf("not able to get data: %v", err)
	}
	for _, v := range r.Data {
		fmt.Println(v.Year, v.Population)
	}
}
