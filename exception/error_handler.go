package exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/wawat7/go-cashier/helper"
	"net/http"
)

func ErrorHandler(c *gin.Context, recovered interface{})  {

	if err, ok := recovered.(validator.ValidationErrors); ok {
		if ok {
			ValidationErrors(c, err)
			return
		}
	}
	if err, ok := recovered.(string); ok {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		return
	}

	if error, ok := recovered.(error); ok {
		c.String(http.StatusBadRequest, error.Error())
		return
	}
	c.AbortWithStatus(http.StatusInternalServerError)
	return
}

func ValidationErrors(c *gin.Context, err error)  {
	response := helper.APIResponse("BAD REQUEST", http.StatusBadRequest, "failed", err.Error())
	c.JSON(http.StatusBadRequest, response)
	return
}