package viacep

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	errors "github.com/BMokarzel/weather-api/pkg/errors"
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

	log.Println("[DEBUG] Response: ", res)

	switch {
	case res.StatusCode < 300:
		var response ViaCepOutput

		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			return ViaCepOutput{}, err
		}
		return response, nil
	case res.StatusCode == 400:
		return ViaCepOutput{}, errors.NewBadRequestError()
	case res.StatusCode == 404:
		return ViaCepOutput{}, errors.NewNotFoundError()
	case res.StatusCode == 422:
		return ViaCepOutput{}, errors.NewUnprocessableEntityError()
	default:
		return ViaCepOutput{}, errors.NewInternalServerError()
	}

}
