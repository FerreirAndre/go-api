package controller

import (
	"fmt"
	"net/http"

	"github.com/FerreirAndre/go-api/src/configuration/logger"
	"github.com/FerreirAndre/go-api/src/configuration/validation"
	"github.com/FerreirAndre/go-api/src/controller/model/request"
	"github.com/FerreirAndre/go-api/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("initializing create user controller",
		zap.String("journey", "create_user"))
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

	logger.Info("user created succesfully",
		zap.String("journey", "create_user"))

	c.JSON(http.StatusOK, response)
}
