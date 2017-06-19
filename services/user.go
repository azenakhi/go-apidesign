package services

import (
	"github.com/azenakhi/go-apidesign/models"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (this UserService) GetUsers() []models.User {
	userModel := models.NewUserModel()
	return userModel.GetUsers()
}
