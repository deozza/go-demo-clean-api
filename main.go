package main

import (
	"api-test/api/handler"
	"api-test/useCases/user"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func main() {
	r := mux.NewRouter()

	userService := user.NewService()
	handler.BuildUserRouter(r, userService)
	log.Fatal(http.ListenAndServe(":8080", r))
}
