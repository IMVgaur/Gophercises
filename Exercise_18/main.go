package main

import (
	"net/http"

	handlers "github.com/Gophercises/Exercise_18/handlers"
)

func main() {
	http.ListenAndServe("localhost:3000", handlers.Handler())
}
