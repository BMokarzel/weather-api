package service

import (
	"context"
	"log"
	"regexp"

	controller_dto "github.com/BMokarzel/weather-api/internal/controller/dto"
)

func (s *Service) GetWeather(ctx context.Context, zipCode string) (interface{}, int) {

	if zipCode == "" {
		return controller_dto.ErrorOutput{
			Message: "can not find zipcode",
		}, 404
	}

	regex := regexp.MustCompile(`^\d{5}-?\d{3}$`)

	if !regex.MatchString(zipCode) {
		log.Printf("Error to parse zipCode. Invalid format.")
		return controller_dto.ErrorOutput{
			Message: "invalid zipcode",
		}, 422
	}

	viaCepRes, err := s.ViaCep.GetLocation(ctx, zipCode)
	if err != nil {
		log.Printf("Error to get location. Error: %s", err)
		return controller_dto.ErrorOutput{
			Message: "invalid zipcode",
		}, 422
	}

	watherRes, err := s.WeatherApi.GetWeather(ctx, viaCepRes.Location)
	if err != nil {
		log.Printf("Error to get weather. Error: %s", err)
		return controller_dto.ErrorOutput{
			Message: "problem to get real time weather. If the problem persists, contact support",
		}, 422
	}

	tempF := watherRes.Weather.CelsiusTemp*1.8 + 32

	tempK := watherRes.Weather.CelsiusTemp + 273

	response := controller_dto.GetWeatherOutput{
		TempC: watherRes.Weather.CelsiusTemp,
		TempF: tempF,
		TempK: tempK,
	}

	return response, 200
}
