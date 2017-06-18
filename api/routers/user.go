package routers

import (
	"github.com/azenakhi/go-learn/restapi/controllers"
	"github.com/gorilla/mux"
)

type UserRouter struct {
}

func NewUserRouter() *UserRouter {
	return &UserRouter{}
}

func (this UserRouter) GetRouter() *mux.Router {
	router := mux.NewRouter()
	userController := controllers.NewUserController()
	router.HandleFunc("/", userController.GetUsers)
	return router
}
