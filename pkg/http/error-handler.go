package http_error

import (
	"encoding/json"
	"errors"
	"net/http"

	pkg_errors "github.com/BMokarzel/weather-api/pkg/errors"
)

type ErrorResponse struct {
	Error string `json:"message"`
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	var (
		statusCode = http.StatusInternalServerError
		message    = "internal server error"
	)

	var badReq *pkg_errors.BadRequestError
	var unprocessable *pkg_errors.UnprocessableEntityError
	var notFound *pkg_errors.NotFount

	switch {
	case errors.As(err, &badReq):
		statusCode = http.StatusBadRequest
		message = badReq.Error()
	case errors.As(err, &unprocessable):
		statusCode = http.StatusUnprocessableEntity
		message = unprocessable.Error()
	case errors.As(err, &notFound):
		statusCode = http.StatusNotFound
		message = notFound.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(ErrorResponse{
		Error: message,
	})
}
