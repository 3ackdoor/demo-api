package entity

type Auditor struct {
	CreatedBy string `gorm:"<-:create"`
	UpdatedBy string `gorm:"<-"`
	DeletedBy string `gorm:"<-:update"`
}
