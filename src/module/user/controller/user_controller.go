package controller

import (
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

func (u *UserController) getAllUsers(c *gin.Context) {
	// r := u.UserService.GetAllUsers()
	c.AbortWithStatusJSON(404, gin.H{
		"error": "404",
	})
	// c.AbortWithError(http.StatusBadRequest, errors.New("AbortWithError"))
		return
	// util.ResponseSuccess(c, r)
}

func (u *UserController) getUserById(c *gin.Context) {
	r := u.UserService.GetUserById(c.Param("id"))
	util.ResponseSuccess(c, r)
}

func (u *UserController) createUser(c *gin.Context) {
	var reqBody dto.UserCreationRequest
	if err := c.BindJSON(&reqBody); err != nil {
		panic(err)
	}

	r := u.UserService.CreateUser(reqBody)
	util.ResponseSuccess(c, r)
}

func (u *UserController) updateUserById(c *gin.Context) {
	var reqBody dto.UserUpdationRequest
	if err := c.BindJSON(&reqBody); err != nil {
		panic(err)
	}

	r := u.UserService.UpdateUserById(c.Param("id"), reqBody)
	util.ResponseSuccess(c, r)
}

func (u *UserController) deleteUserById(c *gin.Context) {
	r := u.UserService.DeleteUserById(c.Param("id"))
	util.ResponseSuccess(c, r)
}
