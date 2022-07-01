package entity

import (
	"github.com/3ackdoor/go-demo-api/src/constant"
	"github.com/3ackdoor/go-demo-api/src/type/null"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt null.Time `gorm:"<-:create;autoCreateTime"`
	UpdatedAt null.Time `gorm:"<-;autoUpdateTime"`
	DeletedAt null.Time `gorm:"<-:update"`
	CreatedBy string    `gorm:"<-:create"`
	UpdatedBy string    `gorm:"<-"`
	DeletedBy string    `gorm:"<-:update"`
}

func (a *BaseModel) BeforeSave(tx *gorm.DB) (err error) {
	a.CreatedBy = constant.DefaultAuditorName
	a.UpdatedBy = constant.DefaultAuditorName
	return
}

func (a *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	a.DeletedBy = constant.DefaultAuditorName
	return
}
