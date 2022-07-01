package service

import (
	"log"

	"github.com/3ackdoor/go-demo-api/src/module/user/converter"
	"github.com/3ackdoor/go-demo-api/src/module/user/dto"
	"github.com/3ackdoor/go-demo-api/src/module/user/entity"
	"github.com/3ackdoor/go-demo-api/src/module/user/repository"
	"github.com/3ackdoor/go-demo-api/src/util"
	"gorm.io/gorm"
)

type UserService interface {
	GetUsers() bool
	GetAllUsers() []dto.UserModel
	GetUserById(string) dto.UserModel
	CreateUser(dto.UserCreationRequest) dto.UserModel
	UpdateUserById(string, dto.UserUpdationRequest) dto.UserModel
	DeleteUserById(string) dto.UserModel
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

func (u *UserServiceImpl) GetAllUsers() []dto.UserModel {
	users := u.UserRepository.FindAll()

	resp := converter.ConvertUserEntitiesToUserModels(users)
	return resp
}

func (u *UserServiceImpl) GetUserById(id string) dto.UserModel {
	idVal := util.ConvertStringToUint(id)
	user := u.UserRepository.FindById(uint(idVal))

	resp := converter.ConvertUserEntityToUserModel(user)
	return resp
}

func (u *UserServiceImpl) CreateUser(request dto.UserCreationRequest) dto.UserModel {
	user := converter.ConvertUserCreationRequestToUserEntity(request)
	u.UserRepository.Save(user)

	resp := converter.ConvertUserEntityToUserModel(user)
	return resp
}

func (u *UserServiceImpl) UpdateUserById(id string, request dto.UserUpdationRequest) dto.UserModel {
	log.Printf("args: %v, %v", id, request)
	idVal := util.ConvertStringToUint(id)
	user := converter.ConvertUserUpdationRequestToUserEntity(request)
	user.ID = idVal
	u.UserRepository.Save(user)

	resp := converter.ConvertUserEntityToUserModel(user)
	return resp
}

func (u *UserServiceImpl) DeleteUserById(id string) dto.UserModel {
	idVal := util.ConvertStringToUint(id)

	user := new(entity.User)
	user.ID = idVal
	u.UserRepository.Delete(user)

	resp := converter.ConvertUserEntityToUserModel(user)
	return resp
}
