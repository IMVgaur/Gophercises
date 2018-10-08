package main

import (
	"net/http"

	handlers "github.com/Gophercises/Exercise_15/handlers"
	middleware "github.com/Gophercises/Exercise_15/middleware"
)

var listenAndServeFunc = http.ListenAndServe

func main() {
	listenAndServeFunc(":3000", middleware.RecoveryMid(handlers.Handler()))
}
