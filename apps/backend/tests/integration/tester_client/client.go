package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	result, err := runRequest()
	if err != nil {
		log.Fatal("error", err)
	}

	fmt.Println(result)
}

func runRequest() (string, error) {
	var err error
	var client = &http.Client{}

	request, err := http.NewRequest("GET", "http://backend:8080/healthcheck", nil)
	if err != nil {
		return "", err
	}

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	s := buf.String()

	return s, nil
}
