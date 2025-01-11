package controller

import (
	"net/http"

	"github.com/FerreirAndre/go-api/src/configuration/logger"
	"github.com/FerreirAndre/go-api/src/configuration/validation"
	"github.com/FerreirAndre/go-api/src/controller/model/request"
	"github.com/FerreirAndre/go-api/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
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

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("user created succesfully",
		zap.String("journey", "create_user"))

	c.JSON(http.StatusOK, "")
}
