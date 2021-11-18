package user

import (
	"github.com/gin-gonic/gin"
	"github.com/wawat7/go-cashier/helper"
	"net/http"
	"strconv"
)

type UserController struct {
	service Service
}

func NewController(service Service) *UserController {
	return &UserController{service: service}
}

func (controller *UserController) Route(app *gin.Engine) {
	route := app.Group("api/users")
	route.GET("/", controller.List)
	route.POST("/", controller.Create)
	route.GET("/:id", controller.GetUser)
	route.PATCH("/:id", controller.Update)
	route.DELETE("/:id", controller.Delete)
}

func (controller *UserController) List(c *gin.Context) {
	users := controller.service.FindAll()
	c.JSON(http.StatusOK, helper.APIResponse("List Users", http.StatusOK, "success", UsersFormat(users)))
	return
}

func (controller *UserController) Create(c *gin.Context) {
	var input CreateUserRequest
	err := c.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	user := User{}
	user = controller.service.Create(input, user)
	c.JSON(http.StatusOK, helper.APIResponse("Success created user", http.StatusOK, "success", UserFormat(user)))
	return
}

func (controller *UserController) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	user := controller.service.FindById(id)
	c.JSON(http.StatusOK, helper.APIResponse("Get User", http.StatusOK, "success", UserFormat(user)))
	return
}

func (controller *UserController) Update(c *gin.Context) {
	var input UpdateUserRequest
	err := c.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	id, _ := strconv.Atoi(c.Params.ByName("id"))
	user := controller.service.FindById(id)
	user.Name = input.Name
	user.PhoneNumber = input.PhoneNumber
	user.Address = input.Address
	user = controller.service.Update(user)
	c.JSON(http.StatusOK, helper.APIResponse("Success update user", http.StatusOK, "success", UserFormat(user)))
	return
}

func (controller *UserController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	controller.service.Delete(id)
	c.JSON(http.StatusOK, helper.APIResponse("Success delete user", http.StatusOK, "success", nil))
	return
}
