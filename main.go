package main

import (
	"net/http"
	"simple-JWT/controllers"
	"simple-JWT/database"
	"simple-JWT/middleware"
)

func main() {

	//connect mongoDB
	database.ConnectMongoDB()

	//handlers
	http.HandleFunc("POST /v1/createuser", middleware.Middleware(controllers.CreateUser))
	http.HandleFunc("GET /v1/user", middleware.Middleware(controllers.UserProfile))

	http.ListenAndServe(":8080", nil)
}
