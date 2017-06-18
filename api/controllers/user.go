package controllers

import (
	"net/http"

	"encoding/json"

	"fmt"

	"github.com/azenakhi/go-learn/restapi/services"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (this UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	userService := services.NewUserService()
	json, _ := json.Marshal(userService.GetUsers())
	fmt.Fprint(w, string(json))
}
