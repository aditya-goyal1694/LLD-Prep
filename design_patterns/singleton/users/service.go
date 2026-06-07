package users

import (
	"lld/design_patterns/singleton/logger"
)

type UserService struct{}

func (u *UserService) PrintUserDetails() {
    logger := logger.GetLogger()
    logger.LogInfo("User id is 1. Name is Aditya.")
}
