package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type config struct {
	ViaCepUrl     string `mapstructure:"VIACEP_API_URL"`
	WeatherApiUrl string `mapstructure:"WEATHER_API_URL"`
	WeatherApiKey string `mapstructure:"WEATHER_API_KEY"`
}

func LoadConfigs() (*config, error) {
	env := os.Getenv("ENV")

	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if env == "local" {
		viper.SetConfigFile("./env")
		viper.SetConfigType("env")
		err := viper.ReadInConfig()
		if err != nil {
			return nil, err
		}
	}

	keys := []string{
		"VIACEP_API_URL",
		"WEATHER_API_URL",
		"WEATHER_API_KEY",
	}

	for _, k := range keys {
		_ = viper.BindEnv(k)
	}

	var cfg config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	cfg.ViaCepUrl = "https://viacep.com.br"
	cfg.WeatherApiKey = "34e6f3d1ecb74dd1941193145251811"
	cfg.WeatherApiUrl = "https://www.weatherapi.com"

	return &cfg, nil
}
