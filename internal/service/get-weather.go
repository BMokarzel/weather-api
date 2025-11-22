package service

import (
	"context"
	"log"
	"regexp"

	controller_dto "github.com/BMokarzel/weather-api/internal/controller/dto"
	errors "github.com/BMokarzel/weather-api/pkg/errors"
)

func (s *Service) GetWeather(ctx context.Context, zipCode string) (interface{}, error) {
	regex := regexp.MustCompile(`^\d{5}-?\d{3}$`)

	if !regex.MatchString(zipCode) {
		log.Printf("Error to parse zipCode. Invalid format.")
		return controller_dto.GetWeatherOutput{}, errors.NewUnprocessableEntityError()
	}

	viaCepRes, err := s.ViaCep.GetLocation(ctx, zipCode)
	if err != nil {
		log.Printf("Error to get location. Error: %s", err)
		return controller_dto.GetWeatherOutput{}, err
	}

	watherRes, err := s.WeatherApi.GetWeather(ctx, viaCepRes.Location)
	if err != nil {
		log.Printf("Error to get weather. Error: %s", err)
		return controller_dto.GetWeatherOutput{}, err
	}

	tempF := watherRes.Current.TempC*1.8 + 32

	tempK := watherRes.Current.TempC + 273

	response := controller_dto.GetWeatherOutput{
		TempC: watherRes.Current.TempC,
		TempF: tempF,
		TempK: tempK,
	}

	return response, nil
}
