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
	token := ""

	if len(args) < 4 {
		log.Fatal("provide the method and the url you want to test")
	}

	if len(args) == 5 {
		token = args[4]
	}

	result, err := runRequest(args[1], args[2], args[3], token)
	if err != nil {
		log.Fatal("error", err)
	}

	fmt.Println(result)
}

func runRequest(method, url, body, token string) (string, error) {
	var err error
	var client = &http.Client{}

	reader := strings.NewReader(body)

	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		return "", err
	}

	if token != "" {
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
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
