package middleware

type BaseDataResponse struct {
	Message string `json:"message,omitempty"`
	Data    any
}

type BasePaginationResponse struct {
	Message string `json:"message,omitempty"`
	Data    any
}

type BaseListResponse struct {
	Message string `json:"message,omitempty"`
	Data    any
}
