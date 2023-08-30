package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func request(url string) {
	if url == "" {
		panic("No url provided")
	}

	// Make request
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	// Check status code
	if res.StatusCode != 200 {
		panic("Status code is not 200")
	}

	// Log response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(body))

}

func main() {
	// Get first arg as URL
	url := os.Args[1]
	sec := 10

	if os.Args != nil && len(os.Args) > 2 {
		secArg := os.Args[2]
		var err error
		sec, err = strconv.Atoi(secArg)
		if err != nil {
			panic(err)
		}
	}

	for {
		request(url)
		// Sleep for 10	seconds
		time.Sleep(time.Duration(sec) * time.Second)
	}
}
