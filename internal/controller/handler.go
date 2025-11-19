package controller

import (
	"encoding/json"
	"net/http"

	"github.com/BMokarzel/weather-api/internal/service"
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

	res, code := h.Service.GetWeather(ctx, zipCode)

	body, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	w.Write(body)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
