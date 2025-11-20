package weather_api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/BMokarzel/weather-api/config"
	"github.com/BMokarzel/weather-api/internal/controller"
	controller_dto "github.com/BMokarzel/weather-api/internal/controller/dto"
	"github.com/BMokarzel/weather-api/internal/service"
	viacep "github.com/BMokarzel/weather-api/pkg/via-cep"
	weatherapi "github.com/BMokarzel/weather-api/pkg/weather-api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	NotFound = "can not found zipcode"
	Invalid  = "invalid zipcode"
)

func TestWeatherApi(t *testing.T) {
	cfg, err := config.LoadConfigs()
	require.NoError(t, err)

	viacep := viacep.New(cfg.ViaCepUrl)
	weather := weatherapi.New(cfg.WeatherApiUrl, cfg.WeatherApiKey)
	service := service.New(viacep, weather)
	handler := controller.New(service)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	q := req.URL.Query()
	q.Add("zipCode", cfg.ZipCode)
	req.URL.RawQuery = q.Encode()

	rec := httptest.NewRecorder()

	handler.GetWeather(rec, req)

	regex := regexp.MustCompile(`^\d{5}-?\d{3}$`)

	if !regex.MatchString(cfg.ZipCode) {
		var output controller_dto.ErrorOutput
		err := json.NewDecoder(rec.Body).Decode(&output)
		require.NoError(t, err)

		assert.Equal(t, Invalid, output.Message)
		return
	}

	switch rec.Code {
	case http.StatusOK:
		var output controller_dto.GetWeatherOutput
		err := json.NewDecoder(rec.Body).Decode(&output)
		require.NoError(t, err)

		assert.NotEmpty(t, output)
	case http.StatusNotFound:
		var output controller_dto.ErrorOutput
		err := json.NewDecoder(rec.Body).Decode(&output)
		require.NoError(t, err)

		assert.Equal(t, NotFound, output.Message)
	case http.StatusUnprocessableEntity:
		var output controller_dto.ErrorOutput
		err := json.NewDecoder(rec.Body).Decode(&output)
		require.NoError(t, err)

		assert.Equal(t, Invalid, output.Message)
	default:
		t.Fatalf("Unexpected status %d. Body: %s", rec.Code, rec.Body.String())
	}
}
