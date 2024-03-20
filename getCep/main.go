package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

var GET_CEP_URL string
var response ViaCepResponse

func main() {
	for _, input := range os.Args[1:] {
		GET_CEP_URL = fmt.Sprintf("https://viacep.com.br/ws/%s/json", input)

		res, err := http.Get(string(GET_CEP_URL))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Request error: %v\n", err)
		}

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Reading response error: %v\n", err)
		}

		defer res.Body.Close()

		err = json.Unmarshal(resBody, &response)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Response parse error: %v\n", err)
		}

		file, err := os.Create("location.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Create file error: %v\n", err)
		}

		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("Postal Code: %s, Location: %s, State: %s\n", response.Cep, response.Localidade, response.Uf))
	}
}
