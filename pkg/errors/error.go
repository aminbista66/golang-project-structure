package errors

import (
	"fmt"
	"net/http"

	"myapp/internal/infrastructure/logger"
	response "myapp/pkg/response"
)

// AppError represents a structured application error that can optionally
// log itself using a provided logger implementation.
type AppError struct {
	Logger  logger.Logger
}


// New creates a new AppError. Pass a logger implementation (may be nil).
func New(l logger.Logger) *AppError {
	return &AppError{Logger: l}
}


// LogRequest logs the provided error along with some request metadata if a
// logger was supplied when the AppError was created.
func (e *AppError) LogRequest(r *http.Request, err error) {
	if e == nil || e.Logger == nil || err == nil || r == nil {
		return
	}
	e.Logger.PrintError(err.Error(), map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	})
}


func (e *AppError) logError(r *http.Request, err error) {
	e.Logger.PrintError(err.Error(), map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	})
}

func (e *AppError) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	data := map[string]interface{}{"error": message}
	err := response.WriteJson(w, status, data)

	if err != nil {
		e.logError(r, err)
		w.WriteHeader(500)
	}
}

func (e *AppError) ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	e.logError(r, err)
	message := "the server encountered a problem and could not process your request"
	e.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (e *AppError) NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	e.errorResponse(w, r, http.StatusNotFound, message)
}

func (e *AppError) MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	e.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (e *AppError) BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	e.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (e *AppError) FailedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	e.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func (e *AppError) EditConflictResponse(w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again"
	e.errorResponse(w, r, http.StatusConflict, message)
}

func (e *AppError) RateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded"
	e.errorResponse(w, r, http.StatusTooManyRequests, message)
}