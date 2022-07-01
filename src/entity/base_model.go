package entity

import (
	"time"

	"github.com/3ackdoor/go-demo-api/src/constant"
	"github.com/3ackdoor/go-demo-api/src/type/null"
	"github.com/3ackdoor/go-demo-api/src/util"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint        `gorm:"primarykey"`
	CreatedAt null.Time   `gorm:"type:timestamp"`
	UpdatedAt null.Time   `gorm:"type:timestamp"`
	DeletedAt null.Time   `gorm:"type:timestamp"`
	CreatedBy null.String `gorm:"<-:create"`
	UpdatedBy null.String `gorm:"<-"`
	DeletedBy null.String `gorm:"<-:update"`
}

func (a *BaseModel) BeforeSave(tx *gorm.DB) (err error) {
	name := constant.DefaultAuditorName
	now := time.Now().UTC()
	a.CreatedAt = null.Time{NullTime: util.NewNullTime(&now)}
	a.UpdatedAt = null.Time{NullTime: util.NewNullTime(&now)}
	a.CreatedBy = null.String{NullString: util.NewNullString(&name)}
	a.UpdatedBy = null.String{NullString: util.NewNullString(&name)}
	return
}

func (a *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	name := constant.DefaultAuditorName
	now := time.Now().UTC()
	a.UpdatedAt = null.Time{NullTime: util.NewNullTime(&now)}
	a.UpdatedBy = null.String{NullString: util.NewNullString(&name)}
	return
}

func (a *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	name := constant.DefaultAuditorName
	now := time.Now().UTC()
	a.DeletedAt = null.Time{NullTime: util.NewNullTime(&now)}
	a.DeletedBy = null.String{NullString: util.NewNullString(&name)}
	return
}
