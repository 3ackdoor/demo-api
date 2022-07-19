package repository

import (
	"errors"

	"github.com/3ackdoor/go-demo-api/src/exception"
	"github.com/3ackdoor/go-demo-api/src/module/user/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() []entity.User
	FindById(id uint) entity.User
	Save(user *entity.User) entity.User
	Update(user *entity.User) entity.User
	Delete(user *entity.User) entity.User
}
type UserRepositoryImpl struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db.Model(&entity.User{}).Session(&gorm.Session{}),
	}
}

func (u *UserRepositoryImpl) FindAll() []entity.User {
	var user []entity.User
	u.Find(&user)
	return user
}

func (u *UserRepositoryImpl) FindById(id uint) entity.User {
	var user entity.User
	res := u.First(&user, "id = ?", id)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		panic(exception.NewValidationException(gorm.ErrRecordNotFound.Error()))
	}

	return user
}

func (u *UserRepositoryImpl) Save(user *entity.User) entity.User {
	u.Statement.Save(&user)
	return *user
}

func (u *UserRepositoryImpl) Update(user *entity.User) entity.User {
	u.Updates(&user)
	return *user
}

func (u *UserRepositoryImpl) Delete(user *entity.User) entity.User {
	u.Statement.Delete(&user)
	return *user
}
