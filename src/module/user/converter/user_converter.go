package converter

import (
	"github.com/3ackdoor/go-demo-api/src/module/user/dto"
	"github.com/3ackdoor/go-demo-api/src/module/user/entity"
)

func ConvertUserCreationToUserEntity(u dto.UserCreationRequest) *entity.User {
	return &entity.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
		Hobby:     u.Hobby,
	}
}

func ConvertUserUpdationRequestToUserEntity(u dto.UserUpdationRequest) *entity.User {
	return &entity.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
	}
}
