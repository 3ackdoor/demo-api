package dto

import (
	"time"

	"github.com/3ackdoor/go-demo-api/src/type/null"
	"gorm.io/gorm"
)

type UserModel struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	CreatedBy string         `json:"createdBy"`
	UpdatedAt time.Time      `json:"updatedAt"`
	UpdatedBy string         `json:"updatedBy"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	DeletedBy string         `json:"deletedBy"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Age       int            `json:"age"`
	Hobby     null.String    `json:"hobby"`
}
