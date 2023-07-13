package api

import (
	"backend/configs"
	"backend/models"
	"encoding/json"
	"log"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	db := configs.Connect_db()
	defer db.Close()

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	password := configs.GeneratePassword()
	user.Password, err = configs.HashPassword(password)
	if err != nil{
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})

}

func Login(w http.ResponseWriter, r *http.Request) {
	db := configs.Connect_db()
	defer db.Close()

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	
	
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})

}
