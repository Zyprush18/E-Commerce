package services

import (
	// "encoding/json"
	"net/http"
)


const (
	Success = http.StatusOK
	Created = http.StatusCreated
	NotFound = http.StatusNotFound
	Forbidden = http.StatusForbidden
	BadRequest = http.StatusBadRequest
	Unauthorized  = http.StatusUnauthorized
	MethodNotAllowed = http.StatusMethodNotAllowed
	InternalServerError = http.StatusInternalServerError
)


type Message struct {
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
	Error any `json:"error,omitempty"`
}


type UserRequest struct {
	Name string `json:"name" validate:"required"` 
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
