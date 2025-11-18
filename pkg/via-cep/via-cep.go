package viacep

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ViaCep struct {
	URL string
}

type ViaCepOutput struct {
	Location string `json:"localidade"`
}

func (v *ViaCep) GetLocation(ctx context.Context, cep string) (ViaCepOutput, error) {

	url := fmt.Sprintf("%s/%s/json/", v.URL, cep)

	res, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return ViaCepOutput{}, err
	}

	if res.Response.StatusCode > 299 {

		return ViaCepOutput{}, nil

	} else {

		var response ViaCepOutput

		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			return ViaCepOutput{}, err
		}

		return response, nil
	}
}
