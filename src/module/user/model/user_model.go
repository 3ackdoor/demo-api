package model

import (
	baseModel "github.com/3ackdoor/go-demo-api/src/model"
	"github.com/3ackdoor/go-demo-api/src/type/null"
)

type UserModel struct {
	baseModel.AuditTrail
	ID        uint        `json:"id"`
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	Age       int         `json:"age"`
	Hobby     null.String `json:"hobby"`
}
