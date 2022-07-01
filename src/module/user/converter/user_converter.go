package converter

import (
	"github.com/3ackdoor/go-demo-api/src/module/user/dto"
	"github.com/3ackdoor/go-demo-api/src/module/user/entity"
)

func ConvertUserCreationRequestToUserEntity(r dto.UserCreationRequest) *entity.User {
	return &entity.User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Age:       r.Age,
		Hobby:     r.Hobby,
	}
}

func ConvertUserUpdationRequestToUserEntity(r dto.UserUpdationRequest, e *entity.User) {
	e.FirstName = r.FirstName
	e.LastName = r.LastName
	e.Age = r.Age
}

func ConvertUserEntityToUserModel(user *entity.User) dto.UserModel {
	return dto.UserModel{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		CreatedBy: user.CreatedBy,
		UpdatedAt: user.UpdatedAt,
		UpdatedBy: user.UpdatedBy,
		DeletedAt: user.DeletedAt,
		DeletedBy: user.DeletedBy,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
		Hobby:     user.Hobby,
	}
}

func ConvertUserEntitiesToUserModels(users *[]entity.User) []dto.UserModel {
	var userModels []dto.UserModel

	for _, user := range *users {
		userModel := ConvertUserEntityToUserModel(&user)
		userModels = append(userModels, userModel)
	}

	return userModels
}
