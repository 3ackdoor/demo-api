package entity

import (
	"github.com/3ackdoor/go-demo-api/src/entity"
	"github.com/3ackdoor/go-demo-api/src/type/null"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	entity.Auditor

	FirstName string
	LastName string
	Age int
	Hobby null.String

}