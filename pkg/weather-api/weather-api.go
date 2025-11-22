package weatherapi

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	errors "github.com/BMokarzel/weather-api/pkg/errors"
)

type WeatherApi struct {
	URL string
	Key string
}

type GetWeatherOutput struct {
	Location Location
	Current  Current `json:"current"`
}

type Current struct {
	TempC float64 `json:"temp_c"`
}

type Location struct {
	Name      string `json:"name"`
	Region    string `json:"region"`
	Country   string `json:"country"`
	LocalTime string `json:"localtime"`
}

type CurrentWeather struct {
	CelsiusTemp float64 `json:"temp_c"`
}

func New(url, key string) *WeatherApi {
	return &WeatherApi{
		URL: url,
		Key: key,
	}
}

func (k *WeatherApi) GetWeather(ctx context.Context, location string) (GetWeatherOutput, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, k.URL, nil)
	if err != nil {
		return GetWeatherOutput{}, err
	}

	c := req.URL.Query()
	c.Add("key", k.Key)
	c.Add("q", location)
	req.URL.RawQuery = c.Encode()

	log.Println("[DEBUG] Request: ", req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return GetWeatherOutput{}, err
	}
	defer res.Body.Close()

	log.Println("[DEBUG] Response: ", res)

	switch {
	case res.StatusCode < 300:
		var response GetWeatherOutput

		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			return GetWeatherOutput{}, err
		}
		return response, nil
	case res.StatusCode == 400:
		return GetWeatherOutput{}, errors.NewBadRequestError()
	case res.StatusCode == 404:
		return GetWeatherOutput{}, errors.NewNotFoundError()
	case res.StatusCode == 422:
		return GetWeatherOutput{}, errors.NewUnprocessableEntityError()
	default:
		return GetWeatherOutput{}, errors.NewInternalServerError()
	}

}
