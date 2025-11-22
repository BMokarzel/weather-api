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
	http_error "github.com/BMokarzel/weather-api/pkg/http"
	viacep "github.com/BMokarzel/weather-api/pkg/via-cep"
	weatherapi "github.com/BMokarzel/weather-api/pkg/weather-api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	BadRequest          = "bad request"
	NotFound            = "can not found zipcode"
	Invalid             = "invalid zipcode"
	InternalServerError = "internal server error"
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
		var output http_error.ErrorResponse
		err := json.NewDecoder(rec.Body).Decode(&output)
		require.NoError(t, err)

		assert.Equal(t, Invalid, output.Error)
		return
	}

	var output controller_dto.GetWeatherOutput
	var outputError http_error.ErrorResponse

	switch rec.Code {
	case http.StatusOK:
		err := json.NewDecoder(rec.Body).Decode(&output)
		require.NoError(t, err)

		assert.NotEmpty(t, output)
	case http.StatusBadRequest:
		err := json.NewDecoder(rec.Body).Decode(&outputError)
		require.NoError(t, err)

		assert.Equal(t, BadRequest, outputError.Error)
	case http.StatusNotFound:
		err := json.NewDecoder(rec.Body).Decode(&outputError)
		require.NoError(t, err)

		assert.Equal(t, NotFound, outputError.Error)
	case http.StatusUnprocessableEntity:
		err := json.NewDecoder(rec.Body).Decode(&output)
		require.NoError(t, err)

		assert.Equal(t, InternalServerError, outputError.Error)
	case http.StatusInternalServerError:

	default:
		t.Fatalf("Unexpected status %d. Body: %s", rec.Code, rec.Body.String())
	}
}
