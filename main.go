package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wawat7/go-cashier/app"
	"github.com/wawat7/go-cashier/exception"
	"github.com/wawat7/go-cashier/helper"
	"github.com/wawat7/go-cashier/services/user"
)

func main() {

	db := app.NewDB()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userController := user.NewController(userService)

	// Setup Gin
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.CustomRecovery(exception.ErrorHandler))

	// Setup Routing
	userController.Route(router)

	// Start App
	err := router.Run(":3000")
	helper.PanicIfError(err)
}
