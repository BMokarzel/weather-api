package main

import (
	"log"
	"net/http"
	"os"

	"github.com/BMokarzel/weather-api/config"
	"github.com/BMokarzel/weather-api/internal/controller"
	"github.com/BMokarzel/weather-api/internal/service"
	viacep "github.com/BMokarzel/weather-api/pkg/via-cep"
	weatherapi "github.com/BMokarzel/weather-api/pkg/weather-api"
	"github.com/go-chi/chi/v5"
)

type Api struct {
	Router  *chi.Mux
	Handler *controller.Handler
}

func main() {
	cfg, err := config.LoadConfigs()
	if err != nil {
		log.Fatalf("Error loading configs. Error %s", err)
		os.Exit(1)
	}

	viaCep := viacep.New(cfg.ViaCepUrl)

	weatherApi := weatherapi.New(cfg.WeatherApiUrl, cfg.WeatherApiKey)

	service := service.New(viaCep, weatherApi)

	handler := controller.New(service)

	router := chi.NewRouter()

	a := Api{
		Router:  router,
		Handler: handler,
	}

	a.Router.Route("/", func(r chi.Router) {
		r.Get("/", a.Handler.GetWeather)
	})

	http.ListenAndServe(":8080", a.Router)
}
