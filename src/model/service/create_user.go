package service

import (
	"fmt"

	"github.com/FerreirAndre/go-api/src/configuration/logger"
	"github.com/FerreirAndre/go-api/src/configuration/rest_err"
	"github.com/FerreirAndre/go-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("init createuser model", zap.String("journey", "createuser"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetName())
	return nil
}
