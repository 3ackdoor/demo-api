package mapper

import (
	baseModel "github.com/3ackdoor/go-demo-api/src/model"
	"github.com/3ackdoor/go-demo-api/src/module/user/dto"
	"github.com/3ackdoor/go-demo-api/src/module/user/entity"
	"github.com/3ackdoor/go-demo-api/src/module/user/model"
)

func MapUserCreationRequestToUserEntity(r dto.UserCreationRequest) *entity.User {
	return &entity.User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Age:       r.Age,
		Hobby:     r.Hobby,
	}
}

func MapUserUpdationRequestToUserEntity(id uint, r dto.UserUpdationRequest, e *entity.User) {
	e.ID = id
	e.FirstName = r.FirstName
	e.LastName = r.LastName
	e.Age = r.Age
}

func MapUserEntityToUserModel(user entity.User) model.UserModel {
	return model.UserModel{
		AuditTrail: baseModel.AuditTrail{
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

func MapUserEntitiesToUserModels(users []entity.User) []model.UserModel {
	var userModels []model.UserModel

	for _, user := range users {
		userModel := MapUserEntityToUserModel(user)
		userModels = append(userModels, userModel)
	}

	return userModels
}
