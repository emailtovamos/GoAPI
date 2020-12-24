package main

import (
	"fmt"
	"github.com/emailtovamos/GoAPI/authentication"
	"github.com/emailtovamos/GoAPI/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", handlers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", handlers.Authenticate).Methods("POST")
	router.HandleFunc("/api/roles", handlers.GetRoles).Methods("GET")

	router.Use(authentication.JwtAuthentication) //attach JWT auth middleware

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
