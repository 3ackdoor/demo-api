package dto

import (
	"encoding/json"
	"log"
)

type ValidationError struct {
	Key         string `json:"key,omitempty"`
	Value       any    `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
}

func (v *ValidationError) BuildMessage() string {
	msgBytes, err := json.Marshal(v)
	if err != nil {
		log.Panic(err)
	}
	return string(msgBytes)
}

type ExternalServiceError struct {
	Description string `json:"description,omitempty"`
}

func (e *ExternalServiceError) BuildMessage() string {
	msgBytes, err := json.Marshal(e)
	if err != nil {
		log.Panic(err)
	}
	return string(msgBytes)
}
