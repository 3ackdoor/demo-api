package entity

import (
	"github.com/3ackdoor/go-demo-api/src/constant"
	"github.com/3ackdoor/go-demo-api/src/type/null"
	"github.com/3ackdoor/go-demo-api/src/util"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint        `gorm:"primarykey"`
	CreatedAt null.Time   `gorm:"<-:create;autoCreateTime"`
	UpdatedAt null.Time   `gorm:"<-;autoUpdateTime"`
	DeletedAt null.Time   `gorm:"<-:update"`
	CreatedBy null.String `gorm:"<-:create"`
	UpdatedBy null.String `gorm:"<-"`
	DeletedBy null.String `gorm:"<-:update"`
}

func (a *BaseModel) BeforeSave(tx *gorm.DB) (err error) {
	name := constant.DefaultAuditorName
	a.CreatedBy = null.String{NullString: util.NewNullString(&name)}
	a.UpdatedBy = null.String{NullString: util.NewNullString(&name)}
	return
}

func (a *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	name := constant.DefaultAuditorName
	a.DeletedBy = null.String{NullString: util.NewNullString(&name)}
	return
}
