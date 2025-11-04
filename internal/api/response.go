package api

import (
	"encoding/json"
	"net/http"
	"solution-for-x/notifychat/internal/app"
)

type successResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type apiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type errorResponse struct {
	Status string    `json:"status"`
	Error  *apiError `json:"error"`
}

type Handler struct {
	app *app.Application
}

func NewHandler(app *app.Application) *Handler {
	return &Handler{
		app: app,
	}
}

func (h *Handler) Success(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
	res := &successResponse{
		Status: "success",
		Data:   data,
	}

	logger := GetLogger(r.Context())
	logger.Info("request successful", "status", statusCode)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) Error(w http.ResponseWriter, r *http.Request, statusCode int, code string, message string) {
	res := &errorResponse{
		Status: "error",
		Error: &apiError{
			Code:    code,
			Message: message,
		},
	}

	logger := GetLogger(r.Context())
	logger.Error(message, "status", statusCode, "code", code)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}
