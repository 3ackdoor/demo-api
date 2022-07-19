package controller

import (
	"github.com/3ackdoor/go-demo-api/src/module/user/dto"
	"github.com/3ackdoor/go-demo-api/src/module/user/service"
	"github.com/3ackdoor/go-demo-api/src/util/response"
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
	response.Success(c, u.UserService.GetAllUsers())
	// r := u.UserService.GetAllUsers()
	// c.AbortWithStatusJSON(404, gin.H{
	// 	"error": "404",
	// })
	// c.AbortWithError(http.StatusBadRequest, errors.New("AbortWithError"))
	// util.ResponseSuccess(c, r)
}

func (u *UserController) getUserById(c *gin.Context) {
	response.Success(c, u.UserService.GetUserById(c.Param("id")))
}

func (u *UserController) createUser(c *gin.Context) {
	var reqBody dto.UserCreationRequest
	if err := c.BindJSON(&reqBody); err != nil {
		panic(err)
	}

	response.Success(c, u.UserService.CreateUser(reqBody))
}

func (u *UserController) updateUserById(c *gin.Context) {
	var reqBody dto.UserUpdationRequest
	if err := c.BindJSON(&reqBody); err != nil {
		panic(err)
	}

	response.Success(c, u.UserService.UpdateUserById(c.Param("id"), reqBody))
}

func (u *UserController) deleteUserById(c *gin.Context) {
	response.Success(c, u.UserService.DeleteUserById(c.Param("id")))
}
