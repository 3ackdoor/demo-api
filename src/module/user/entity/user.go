package entity

import (
	"github.com/3ackdoor/go-demo-api/src/constant"
	"github.com/3ackdoor/go-demo-api/src/entity"
	"github.com/3ackdoor/go-demo-api/src/type/null"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	entity.Auditor

	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Age       int    `gorm:"not null"`
	Hobby     null.String
}

func (User) TableName() string {
	return "mst_user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedBy = constant.DefaultAuditorName
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedBy = constant.DefaultAuditorName
	return nil
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	u.DeletedBy = constant.DefaultAuditorName
	return nil
}
