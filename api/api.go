package api

import (
	"net/http"

	"log"

	"github.com/azenakhi/go-apidesign/api/routers"
)

func Run() {
	userRouter := routers.NewUserRouter()
	log.Fatal(http.ListenAndServe(":8080", userRouter.GetRouter()))
}
