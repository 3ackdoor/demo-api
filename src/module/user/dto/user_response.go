package dto

import (
	"github.com/3ackdoor/go-demo-api/src/dto"
	"github.com/3ackdoor/go-demo-api/src/type/null"
)

type UserModel struct {
	dto.AuditTrail
	ID        uint        `json:"id"`
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	Age       int         `json:"age"`
	Hobby     null.String `json:"hobby"`
}
