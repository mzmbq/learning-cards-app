package apiserver

import (
	"fmt"
	"log"
	"net/http"
)

type APIError struct {
	StatusCode int `json:"statusCode"`
	Msg        any `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("API error: %v", e.Msg)
}

func NewAPIError(statusCode int, msg string) APIError {
	return APIError{
		StatusCode: statusCode,
		Msg:        msg,
	}
}

func InvalidRequestData(errors map[string]string) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Msg:        errors,
	}
}

func InvalidJSON() APIError {
	return NewAPIError(http.StatusBadRequest, "invalid JSON request data")
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func MakeHandler(h APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			apiErr, ok := err.(APIError)
			if !ok {
				apiErr = NewAPIError(
					http.StatusInternalServerError,
					"internal server error",
				)
			}
			WriteJSON(w, apiErr.StatusCode, apiErr)
			log.Println("HTTP API error:", err.Error(), r.URL.Path)
		}
	}
}
