package controller

import (
	"fmt"

	"github.com/FerreirAndre/go-api/src/configuration/rest_err"
	"github.com/FerreirAndre/go-api/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil{
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("there are some incorrect fields, errors = %s", err),
			)
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}
