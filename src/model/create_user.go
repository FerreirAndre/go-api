package model

import (
	"fmt"

	"github.com/FerreirAndre/go-api/src/configuration/logger"
	"github.com/FerreirAndre/go-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("init createuser model", zap.String("journey", "createuser"))

	ud.EncryptPassword()

	fmt.Println(ud)
	return nil
}
