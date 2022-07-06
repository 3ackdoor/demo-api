package mapper

import (
	"github.com/3ackdoor/go-demo-api/src/dto"
	userDto "github.com/3ackdoor/go-demo-api/src/module/user/dto"
	"github.com/3ackdoor/go-demo-api/src/module/user/entity"
)

func MapUserCreationRequestToUserEntity(r userDto.UserCreationRequest) *entity.User {
	return &entity.User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Age:       r.Age,
		Hobby:     r.Hobby,
	}
}

func MapUserUpdationRequestToUserEntity(id uint, r userDto.UserUpdationRequest, e *entity.User) {
	e.ID = id
	e.FirstName = r.FirstName
	e.LastName = r.LastName
	e.Age = r.Age
}

func MapUserEntityToUserModel(user entity.User) userDto.UserModel {
	return userDto.UserModel{
		AuditTrail: dto.AuditTrail{
			CreatedBy: user.CreatedBy,
			UpdatedBy: user.UpdatedBy,
			DeletedBy: user.DeletedBy,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},

		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
		Hobby:     user.Hobby,
	}
}

func MapUserEntitiesToUserModels(users []entity.User) []userDto.UserModel {
	var userModels []userDto.UserModel

	for _, user := range users {
		userModel := MapUserEntityToUserModel(user)
		userModels = append(userModels, userModel)
	}

	return userModels
}
