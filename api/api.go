package api

import (
	"net/http"

	"log"

	"github.com/azenakhi/go-learn/restapi/routers"
)

func Run() {
	userRouter := routers.NewUserRouter()
	log.Fatal(http.ListenAndServe(":8080", userRouter.GetRouter()))
}
