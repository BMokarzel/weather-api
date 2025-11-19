package viacep

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ViaCep struct {
	URL string
}

type ViaCepOutput struct {
	Location string `json:"localidade"`
}

func New(url string) *ViaCep {
	return &ViaCep{
		URL: url,
	}
}

func (v *ViaCep) GetLocation(ctx context.Context, cep string) (ViaCepOutput, error) {

	url := fmt.Sprintf("%s/%s/json/", v.URL, cep)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return ViaCepOutput{}, err
	}

	log.Println("[DEBUG] Request: ", req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ViaCepOutput{}, err
	}

	if res.StatusCode > 299 {

		log.Println("[DEBUG] Response: ", res)

		return ViaCepOutput{}, fmt.Errorf("")

	} else {

		var response ViaCepOutput

		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			return ViaCepOutput{}, err
		}

		log.Println("[DEBUG] Response: ", res)

		return response, nil
	}
}
