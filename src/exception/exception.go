package exception

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

const defaultErrorMessage string = "Something went wrong"

const (
	InternalServiceExceptionCode Code = "INTERNAL_SERVICE_ERROR"
	InputValidationExceptionCode Code = "VALIDATION_INPUT_ERROR"
	ValidationExceptionCode      Code = "VALIDATION_ERROR"
	ExternalServiceExceptionCode Code = "EXTERNAL_SERVICE_ERROR"
)

var statusCode = map[Code]int{
	InternalServiceExceptionCode: http.StatusInternalServerError,
	InputValidationExceptionCode: http.StatusBadRequest,
	ValidationExceptionCode:      http.StatusBadRequest,
	ExternalServiceExceptionCode: http.StatusServiceUnavailable,
}

// StatusCode returns HTTP status code for the exception code. It returns int(0) if the exception code is unknown.
func StatusCode(text Code) int {
	return statusCode[text]
}

type Code string

type ErrorDetail struct {
	Description string `json:"description,omitempty"`
}

type ValidationErrorDetail struct {
	Key         string
	Description string `json:"description,omitempty"`
	Value       any
}

type InternalServiceException struct {
	Code    `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Detail  ErrorDetail `json:"detail,omitempty"`
}

func NewInternalServiceException(description ...string) *InternalServiceException {
	e := new(InternalServiceException)
	e.Build(description...)
	return e
}

func (*InternalServiceException) StatusCode() int {
	return StatusCode(InternalServiceExceptionCode)
}

func (i *InternalServiceException) Build(description ...string) {
	i.Code = InternalServiceExceptionCode
	i.Message = http.StatusText(i.StatusCode())

	if len(description) == 1 && len(description[0]) != 0 {
		i.Detail = ErrorDetail{
			Description: description[0],
		}
	} else {
		i.Detail = ErrorDetail{
			Description: defaultErrorMessage,
		}
	}
}

type InputValidationException struct {
	Code    `json:"code,omitempty"`
	Message string                  `json:"message,omitempty"`
	Details []ValidationErrorDetail `json:"details,omitempty"`
}

func NewInputValidationException(fieldErrors validator.ValidationErrors) *InputValidationException {
	e := new(InputValidationException)
	e.Build(fieldErrors)
	return e
}

func (*InputValidationException) StatusCode() int {
	return StatusCode(InputValidationExceptionCode)
}

func (i *InputValidationException) Build(fieldErrors validator.ValidationErrors) {
	i.Code = InputValidationExceptionCode
	i.Message = http.StatusText(i.StatusCode())
	if len(fieldErrors) > 0 {
		for _, fieldError := range fieldErrors {
			i.Details = append(i.Details, ValidationErrorDetail{
				Key:         fieldError.Field(),
				Value:       fieldError.Value(),
				Description: fieldError.Error(),
			})
		}
	} else {
		panic(NewInternalServiceException())
	}
}

type ValidationException struct {
	Code    `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Detail  ErrorDetail `json:"detail,omitempty"`
}

func NewValidationException(description string) *ValidationException {
	e := new(ValidationException)
	e.Build(description)
	return e
}


func (*ValidationException) StatusCode() int {
	return StatusCode(ValidationExceptionCode)
}

func (v *ValidationException) Build(description string) {
	v.Code = ValidationExceptionCode
	v.Message = http.StatusText(v.StatusCode())
	v.Detail = ErrorDetail{
		Description: description,
	}
}

type ExternalServiceException struct {
	Code    `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Detail  ErrorDetail `json:"detail,omitempty"`
}

func NewExternalServiceException(description string) *ExternalServiceException {
	e := new(ExternalServiceException)
	e.Build(description)
	return e
}

func (*ExternalServiceException) StatusCode() int {
	return StatusCode(ExternalServiceExceptionCode)
}

func (e *ExternalServiceException) Build(description string) {
	e.Code = ExternalServiceExceptionCode
	e.Message = http.StatusText(e.StatusCode())
	e.Detail = ErrorDetail{
		Description: description,
	}
}
