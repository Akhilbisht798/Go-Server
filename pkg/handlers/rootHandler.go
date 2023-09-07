package handlers

import (
	"booking-app/pkg/db"
	"booking-app/pkg/models"
	"encoding/json"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request Body", http.StatusBadRequest)
		return
	}

	if err := db.CreateUser(user); err != nil {
		http.Error(w, "Error Creating User", http.StatusInternalServerError)
	}

	w.Write([]byte("Succesfully created the User"))
}