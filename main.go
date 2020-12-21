package main

import (
	"github.com/emailtovamos/GoAPI/authentication"
	"github.com/emailtovamos/GoAPI/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"fmt"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", handlers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", handlers.Authenticate).Methods("POST")
	router.HandleFunc("/api/roles", handlers.GetRoles).Methods("GET")
	//router.HandleFunc("/api/me/contacts", handlers.GetContactsFor).Methods("GET") //  user/2/contacts

	router.Use(authentication.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}