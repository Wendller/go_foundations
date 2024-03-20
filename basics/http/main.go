package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GHUserResponse struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	URL       string `json:"url"`
	ReposURL  string `json:"repos_url"`
	Location  string `json:"location"`
}

func main() {
	res, err := http.Get("https://api.github.com/users/Wendller")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var response GHUserResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		panic(err)
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))

}
