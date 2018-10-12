package main

import (
	"net/http"

	handlers "github.com/IMVgaur/Gophercises/Exercise_15/handlers"
	middleware "github.com/IMVgaur/Gophercises/Exercise_15/middleware"
)

var listenAndServeFunc = http.ListenAndServe

func main() {
	listenAndServeFunc(":3000", middleware.RecoveryMid(handlers.Handler()))
}
