package converter

import (
	"github.com/3ackdoor/go-demo-api/src/dto"
	userDto "github.com/3ackdoor/go-demo-api/src/module/user/dto"
	"github.com/3ackdoor/go-demo-api/src/module/user/entity"
)

func ConvertUserCreationRequestToUserEntity(r userDto.UserCreationRequest) *entity.User {
	return &entity.User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Age:       r.Age,
		Hobby:     r.Hobby,
	}
}

func ConvertUserUpdationRequestToUserEntity(r userDto.UserUpdationRequest, e *entity.User) {
	e.FirstName = r.FirstName
	e.LastName = r.LastName
	e.Age = r.Age
}

func ConvertUserEntityToUserModel(user *entity.User) userDto.UserModel {
	return userDto.UserModel{
		BaseResponse: dto.BaseResponse{
			ID:        user.ID,
			CreatedBy: user.CreatedBy,
			UpdatedBy: user.UpdatedBy,
			DeletedBy: user.DeletedBy,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},

		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
		Hobby:     user.Hobby,
	}
}

func ConvertUserEntitiesToUserModels(users *[]entity.User) []userDto.UserModel {
	var userModels []userDto.UserModel

	for _, user := range *users {
		userModel := ConvertUserEntityToUserModel(&user)
		userModels = append(userModels, userModel)
	}

	return userModels
}
