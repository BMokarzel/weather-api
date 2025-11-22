package controller

import (
	"encoding/json"
	"net/http"

	"github.com/BMokarzel/weather-api/internal/service"
	errors "github.com/BMokarzel/weather-api/pkg/errors"
	http_error "github.com/BMokarzel/weather-api/pkg/http"
)

type Handler struct {
	Service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetWeather(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	zipCode := r.URL.Query().Get("zipCode")

	res, err := h.Service.GetWeather(ctx, zipCode)
	if err != nil {
		http_error.ErrorHandler(w, r, err)
		return
	}

	body, err := json.Marshal(res)
	if err != nil {
		http_error.ErrorHandler(w, r, errors.NewInternalServerError())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
