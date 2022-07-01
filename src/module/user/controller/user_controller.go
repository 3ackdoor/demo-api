package controller

import (
	"log"

	"github.com/3ackdoor/go-demo-api/src/module/user/dto"
	"github.com/3ackdoor/go-demo-api/src/module/user/service"
	"github.com/3ackdoor/go-demo-api/src/util"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(rg *gin.RouterGroup, us service.UserService) {

	uc := &UserController{
		UserService: us,
	}

	userCtrl := rg.Group("/users")
	{
		userCtrl.GET("/", uc.getAllUsers)
		userCtrl.GET("/:id/", uc.getUserById)
		userCtrl.POST("/", uc.createUser)
		userCtrl.PUT("/:id/", uc.updateUserById)
		userCtrl.DELETE("/:id/", uc.deleteUserById)
	}
}

func (u *UserController) getAllUsers(ctx *gin.Context) {
	r := u.UserService.GetAllUsers()
	util.ResponseSuccess(ctx, r)
}

func (u *UserController) getUserById(ctx *gin.Context) {
	r := u.UserService.GetUserById(ctx.Param("id"))
	util.ResponseSuccess(ctx, r)
}

func (u *UserController) createUser(ctx *gin.Context) {
	var reqBody dto.UserCreationRequest
	if err := ctx.Bind(&reqBody); err != nil {
		log.Fatal(err)
	}

	r := u.UserService.CreateUser(reqBody)
	util.ResponseSuccess(ctx, r)
}

func (u *UserController) updateUserById(ctx *gin.Context) {
	var reqBody dto.UserUpdationRequest
	if err := ctx.Bind(&reqBody); err != nil {
		log.Fatal(err)
	}

	r := u.UserService.UpdateUserById(ctx.Param("id"), reqBody)
	util.ResponseSuccess(ctx, r)
}

func (u *UserController) deleteUserById(ctx *gin.Context) {
	r := u.UserService.DeleteUserById(ctx.Param("id"))
	util.ResponseSuccess(ctx, r)
}
