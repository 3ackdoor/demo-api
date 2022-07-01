package dto

import (
	"github.com/3ackdoor/go-demo-api/src/module/user/entity"
	"github.com/3ackdoor/go-demo-api/src/type/null"
)

type UserCreationRequest struct {
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	Age       int         `json:"age"`
	Hobby     null.String `json:"hobby"`
}

func (u UserCreationRequest) ToUserEntiy() *entity.User {
	return &entity.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
		Hobby:     u.Hobby,
	}
}

type UserUpdationRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}
