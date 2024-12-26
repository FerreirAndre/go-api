package controller

import (
	"fmt"

	"github.com/FerreirAndre/go-api/src/configuration/validation"
	"github.com/FerreirAndre/go-api/src/controller/model/request"
	"github.com/FerreirAndre/go-api/src/controller/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}
	fmt.Println(userRequest)

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	fmt.Println(response)
}
