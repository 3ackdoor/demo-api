package service

import (
	"github.com/3ackdoor/go-demo-api/src/exception"
	"github.com/3ackdoor/go-demo-api/src/module/user/dto"
	"github.com/3ackdoor/go-demo-api/src/module/user/entity"
	"github.com/3ackdoor/go-demo-api/src/module/user/mapper"
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
	panic(exception.NewInternalServiceException("rise error"))

	resp := mapper.MapUserEntitiesToUserModels(users)
	return resp
}

func (u *UserServiceImpl) GetUserById(id string) dto.UserModel {
	idVal := util.ConvertStringToUint(id)
	user := u.UserRepository.FindById(uint(idVal))

	resp := mapper.MapUserEntityToUserModel(user)
	return resp
}

func (u *UserServiceImpl) CreateUser(request dto.UserCreationRequest) dto.UserModel {
	user := mapper.MapUserCreationRequestToUserEntity(request)
	u.UserRepository.Save(user)

	resp := mapper.MapUserEntityToUserModel(*user)
	return resp
}

func (u *UserServiceImpl) UpdateUserById(id string, request dto.UserUpdationRequest) dto.UserModel {
	idVal := util.ConvertStringToUint(id)

	user := u.UserRepository.FindById(idVal)

	mapper.MapUserUpdationRequestToUserEntity(idVal, request, &user)
	u.UserRepository.Update(&user)

	resp := mapper.MapUserEntityToUserModel(user)
	return resp
}

func (u *UserServiceImpl) DeleteUserById(id string) dto.UserModel {
	idVal := util.ConvertStringToUint(id)

	user := new(entity.User)
	user.ID = idVal
	u.UserRepository.Delete(user)

	resp := mapper.MapUserEntityToUserModel(*user)
	return resp
}
