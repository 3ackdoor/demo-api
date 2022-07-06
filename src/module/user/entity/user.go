package entity

import (
	"github.com/3ackdoor/go-demo-api/src/entity"
	"github.com/3ackdoor/go-demo-api/src/type/null"
)

type User struct {
	entity.AuditTrail
	ID        uint   `gorm:"primarykey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Age       int    `gorm:"not null"`
	Hobby     null.String
}

func (User) TableName() string {
	return "mst_user"
}
