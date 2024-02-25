package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 4 {
		log.Fatal("provide the method and the url you want to test")
	}

	result, err := runRequest(args[1], args[2], args[3])
	if err != nil {
		log.Fatal("error", err)
	}

	fmt.Println(result)
}

func runRequest(method string, url string, body string) (string, error) {
	var err error
	var client = &http.Client{}

	reader := strings.NewReader(body)

	request, err := http.NewRequest(method, url, reader)
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
