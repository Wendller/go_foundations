package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	client := http.Client{Timeout: time.Second}

	res, err := client.Get("https://api.github.com/users/Wendller")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))
}
