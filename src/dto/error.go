package dto

type ValidationError struct {
	Key string
	Value any
	Description error
}

type ExternalServiceError struct {
	Description error
}