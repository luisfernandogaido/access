package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	const n = 128
	const c = 32
	sem := make(chan struct{}, c)
	for i := 0; i < n; i++ {
		sem <- struct{}{}
		go func() {
			defer func() { <-sem }()
			lenBytes, err := list()
			if err != nil {
				log.Println(err)
			} else {
				fmt.Println(lenBytes)
			}
		}()
	}
	for i := 0; i < c; i++ {
		sem <- struct{}{}
	}
}

func list() (int, error) {
	res, err := http.Get("http://localhost:4013/?" +
		"ini=2020-07-10T00:00:00&" +
		"fim=2020-07-10T23:59:59&" +
		"app=profinanc&" +
		"token=yjZFwp3d5ww1h4ja6Uya",
	)
	if err != nil {
		return 0, fmt.Errorf("list: %w", err)
	}
	if res.StatusCode != 200 {
		return 0, fmt.Errorf("list: status code %v", res.StatusCode)
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	return len(b), nil
}
