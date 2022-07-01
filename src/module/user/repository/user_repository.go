package repository

import (
	"github.com/3ackdoor/go-demo-api/src/module/user/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() *[]entity.User
}
type UserRepositoryImpl struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db.Model(&entity.User{}),
	}
}

func (u *UserRepositoryImpl) FindAll() *[]entity.User{
	var user []entity.User
	u.Find(&user)

	return &user
}