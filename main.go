package main

import (
	"./controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/v1/user/sign-up", controllers.SignUp).Methods("POST")
	router.HandleFunc("/v1/user/sign-in", controllers.SignIn).Methods("POST")

	port := os.Getenv("port")

	if port == "" {
		port = "8080"
	}

	fmt.Print(port)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Println(err)
	}
}
