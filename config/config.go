package config

import (
	"github.com/spf13/viper"
)

type config struct {
	ViaCepUrl     string `mapstructure:"VIACEP_API_URL"`
	WeatherApiUrl string `mapstructure:"WEATHER_API_URL"`
	WeatherApiKey string `mapstructure:"WEATHER_API_KEY"`
	ZipCode       string `mapstructure:"ZIPCODE"`
}

func LoadConfigs() (*config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
