package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"simple-JWT/database"
	"simple-JWT/models"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Initialize user model
	var user models.User

	// Parse JSON into user model
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	//default value
	user.CreatedAt = time.Now()

	// Initialize validator
	validate := validator.New()

	// Validate user struct
	if err := validate.Struct(user); err != nil {
		// If validation fails, return detailed error
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Validation failed",
			"details": err.Error(),
		})
		return
	}
	//mongodb
	collection := database.GetCollection()
	filter := bson.D{{Key: "email", Value: user.Email}}

	var foundUser models.User
	findErr := collection.FindOne(context.TODO(), filter).Decode(&foundUser)

	if findErr == mongo.ErrNoDocuments {
		//storing user in db
		if _, err := collection.InsertOne(context.TODO(), user); err != nil {
			http.Error(w, "error while inserting data", http.StatusInternalServerError)
			log.Print(err)
			return
		}
		// Success response
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Your account created successfully.\n JWT token : %s", CreateToken(user.FirstName, user.LastName, user.Email))
		return
	}

	//user already exist
	w.WriteHeader(http.StatusConflict)
	json.NewEncoder(w).Encode(map[string]string{
		"error":   "User already exists.",
		"message": fmt.Sprintf("User with email %s already exists.", user.Email),
	})
}

func UserProfile(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		fmt.Fprintf(w, "missing Authorization Header.")
		return
	}
	err := VerifyToken(authHeader)
	if err != nil {
		// token not valid
		fmt.Fprintf(w, "You are not Authorized user.")
		return
	}
	// valid token
	fmt.Fprintf(w, "You are Authorized User.")
}
