package weatherapi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	pkg_errors "github.com/BMokarzel/weather-api/pkg/errors"
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

func New(url, key string) *WeatherApi {
	return &WeatherApi{
		URL: url,
		Key: key,
	}
}

func (k *WeatherApi) GetWeather(ctx context.Context, location string) (GetWeatherOutput, error) {

	url := fmt.Sprintf("%s/v1/current.json", k.URL)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
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

	if res.StatusCode == 422 {
		log.Println("[DEBUG] Response: ", res)

		return GetWeatherOutput{}, pkg_errors.NewNotFoundError()

	} else if res.StatusCode > 299 && res.StatusCode != 404 {
		log.Println("[DEBUG] Response: ", res)

		return GetWeatherOutput{}, pkg_errors.NewNotFoundError()
	} else {
		log.Println("[DEBUG] Response: ", res)

		var response GetWeatherOutput

		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			return GetWeatherOutput{}, err

		}

		return response, nil
	}

}
