package main

import (
	"fmt"
	"net/http"
)

func request() {
	res, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}

	if res.StatusCode != 200 {
		panic("Status code is not 200")
	}

	fmt.Println(res)
}

func main() {
	for {
		request()
	}
}
