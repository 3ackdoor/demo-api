package dto

type ValidationError struct {
	Key         string `json:"key"`
	Value       any    `json:"value"`
	Description string `json:"description"`
}

type ExternalServiceError struct {
	Description string `json:"description"`
}
