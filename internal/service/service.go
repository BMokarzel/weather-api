package service

import (
	viacep "github.com/BMokarzel/weather-api/pkg/via-cep"
	weatherapi "github.com/BMokarzel/weather-api/pkg/weather-api"
)

type Service struct {
	ViaCep     *viacep.ViaCep
	WeatherApi *weatherapi.WeatherApi
}

func New(viaCep *viacep.ViaCep, weatherApi *weatherapi.WeatherApi) *Service {
	return &Service{
		ViaCep:     viaCep,
		WeatherApi: weatherApi,
	}
}
