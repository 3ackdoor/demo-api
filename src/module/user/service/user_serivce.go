package service

import (
	"log"

	"github.com/3ackdoor/go-demo-api/src/module/user/dto"
	"github.com/3ackdoor/go-demo-api/src/module/user/repository"
	"gorm.io/gorm"
)

type UserService interface {
	GetUsers() bool
	GetAllUsers() bool
	GetUserById(any) bool
	CreateUser(dto.UserCreationRequest) bool
	UpdateUserById(any, dto.UserUpdationRequest) bool
	DeleteUserById(any) bool
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(db *gorm.DB) *UserServiceImpl {

	repo := repository.NewUserRepository(db)

	return &UserServiceImpl{
		UserRepository: repo,
	}
}

func (u *UserServiceImpl) GetUsers() bool {
	return true
}

func (u *UserServiceImpl) GetAllUsers() bool {
	users := u.UserRepository.FindAll()
	log.Printf("users: %v", users)
	return true
}

func (u *UserServiceImpl) GetUserById(id any) bool {
	log.Printf("args: %v", id)
	return true
}

func (u *UserServiceImpl) CreateUser(request dto.UserCreationRequest) bool {
	log.Printf("args: %v", request)
	return true
}

func (u *UserServiceImpl) UpdateUserById(id any, request dto.UserUpdationRequest) bool {
	log.Printf("args: %v, %v", id, request)
	return true
}

func (u *UserServiceImpl) DeleteUserById(id any) bool {
	log.Printf("args: %v", id)
	return true
}
