package weatherapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherApi struct {
	URL string
	Key string
}

type GetWeatherOutput struct {
	Location Location
	Weather  CurrentWeather
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

func (k *WeatherApi) GetWeather(ctx context.Context, location string) (GetWeatherOutput, error) {

	url := fmt.Sprintf("%s/current.json", k.URL)

	var req *http.Request

	req.URL.Query().Add("key", k.Key)
	req.URL.Query().Add("q", location)

	res, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return GetWeatherOutput{}, err
	}

	var response GetWeatherOutput

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return GetWeatherOutput{}, err
	}

	return GetWeatherOutput{}, nil
}
