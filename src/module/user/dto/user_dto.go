package dto

import (
	"github.com/3ackdoor/go-demo-api/src/type/null"
)

type UserCreationRequest struct {
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	Age       int         `json:"age"`
	Hobby     null.String `json:"hobby"`
}

type UserUpdationRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}
