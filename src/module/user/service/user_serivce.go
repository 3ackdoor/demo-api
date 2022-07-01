package service

import (
	"log"

	"github.com/3ackdoor/go-demo-api/src/module/user/dto"
	"github.com/3ackdoor/go-demo-api/src/module/user/repository"
	"github.com/3ackdoor/go-demo-api/src/util"
	"gorm.io/gorm"
)

type UserService interface {
	GetUsers() bool
	GetAllUsers() bool
	GetUserById(string) bool
	CreateUser(dto.UserCreationRequest) bool
	UpdateUserById(string, dto.UserUpdationRequest) bool
	DeleteUserById(string) bool
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

func (u *UserServiceImpl) GetUserById(id string) bool {
	log.Printf("args: %v", id)

	idVal := util.ConvertStringToUint(id)
	users := u.UserRepository.FindById(uint(idVal))

	log.Printf("users: %v", users)

	return true
}

func (u *UserServiceImpl) CreateUser(request dto.UserCreationRequest) bool {
	log.Printf("args: %v", request)
	return true
}

func (u *UserServiceImpl) UpdateUserById(id string, request dto.UserUpdationRequest) bool {
	log.Printf("args: %v, %v", id, request)
	return true
}

func (u *UserServiceImpl) DeleteUserById(id string) bool {
	log.Printf("args: %v", id)
	return true
}
