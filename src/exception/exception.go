package exception

import (
	"encoding/json"
	"log"
)

const (
	ExternalServiceExceptionName = "ExternalService"
	ValidationExceptionName      = "ValidationFailed"
)

type Code string

type ValidationException struct {
	Code        `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

func (v *ValidationException) BuildMessage() string {
	v.Code = ValidationExceptionName
	msgBytes, err := json.Marshal(v)
	if err != nil {
		log.Panic(err)
	}
	return string(msgBytes)
}

type ExternalServiceException struct {
	Code
	Description string `json:"description,omitempty"`
}

func (e *ExternalServiceException) BuildMessage() string {
	e.Code = ExternalServiceExceptionName
	msgBytes, err := json.Marshal(e)
	if err != nil {
		log.Panic(err)
	}
	return string(msgBytes)
}
