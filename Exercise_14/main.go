package main

import (
	"fmt"
	"log"
	"net/http"
)

//main function
func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", welcome)
	server.HandleFunc("/panic/", panicDemo)
	server.HandleFunc("/afterPanic/", afterPanicDemo)
	log.Fatal(http.ListenAndServe(":3000", PanicHandler(server)))
}

//PanicHandler ... handle the situation whenever panic occures while execution
//input : http handler
//return : http handler function
func PanicHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println("This is the error : ", err)
				http.Error(w, err.(string), http.StatusInternalServerError)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	}
}

//welcome function displays welcome note on homepage
func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
	w.WriteHeader(http.StatusInternalServerError)
}

//custom panic
func panicDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before panic")
	createPanic()
}

//this function reveals the situantions after panic has be handled
func afterPanicDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("panic inside afterPanicDemo function")
	createPanic()
}

//panics here
func createPanic() {
	panic("Panic says : Statue")
}
