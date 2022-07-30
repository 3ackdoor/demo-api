package service

import (
	"github.com/3ackdoor/go-demo-api/src/exception"
	"github.com/3ackdoor/go-demo-api/src/module/user/dto"
	"github.com/3ackdoor/go-demo-api/src/module/user/entity"
	"github.com/3ackdoor/go-demo-api/src/module/user/mapper"
	"github.com/3ackdoor/go-demo-api/src/module/user/model"
	"github.com/3ackdoor/go-demo-api/src/module/user/repository"
	"github.com/3ackdoor/go-demo-api/src/util"
	"gorm.io/gorm"
)

type UserService interface {
	GetUsers() bool
	GetAllUsers() []model.UserModel
	GetUserById(string) model.UserModel
	CreateUser(dto.UserCreationRequest) model.UserModel
	UpdateUserById(string, dto.UserUpdationRequest) model.UserModel
	DeleteUserById(string) model.UserModel
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

func (u *UserServiceImpl) GetAllUsers() []model.UserModel {
	users := u.UserRepository.FindAll()
	panic(exception.NewValidationException("rise error"))

	resp := mapper.MapUserEntitiesToUserModels(users)
	return resp
}

func (u *UserServiceImpl) GetUserById(id string) model.UserModel {
	idVal := util.ConvertStringToUint(id)
	user := u.UserRepository.FindById(uint(idVal))

	resp := mapper.MapUserEntityToUserModel(user)
	return resp
}

func (u *UserServiceImpl) CreateUser(request dto.UserCreationRequest) model.UserModel {
	user := mapper.MapUserCreationRequestToUserEntity(request)
	u.UserRepository.Save(user)

	resp := mapper.MapUserEntityToUserModel(*user)
	return resp
}

func (u *UserServiceImpl) UpdateUserById(id string, request dto.UserUpdationRequest) model.UserModel {
	idVal := util.ConvertStringToUint(id)

	user := u.UserRepository.FindById(idVal)

	mapper.MapUserUpdationRequestToUserEntity(idVal, request, &user)
	u.UserRepository.Update(&user)

	resp := mapper.MapUserEntityToUserModel(user)
	return resp
}

func (u *UserServiceImpl) DeleteUserById(id string) model.UserModel {
	idVal := util.ConvertStringToUint(id)

	user := new(entity.User)
	user.ID = idVal
	u.UserRepository.Delete(user)

	resp := mapper.MapUserEntityToUserModel(*user)
	return resp
}
