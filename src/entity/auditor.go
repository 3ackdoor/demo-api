package entity

import (
	"github.com/3ackdoor/go-demo-api/src/constant"
	"gorm.io/gorm"
)

type Auditor struct {
	CreatedBy string `gorm:"<-:create"`
	UpdatedBy string `gorm:"<-"`
	DeletedBy string `gorm:"<-:update"`
}

func (a *Auditor) BeforeCreate(tx *gorm.DB) (err error) {
	a.CreatedBy = constant.DefaultAuditorName
	return nil
}

func (a *Auditor) BeforeUpdate(tx *gorm.DB) (err error) {
	a.UpdatedBy = constant.DefaultAuditorName
	return nil
}

func (a *Auditor) BeforeDelete(tx *gorm.DB) (err error) {
	a.DeletedBy = constant.DefaultAuditorName
	return nil
}
