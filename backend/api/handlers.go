package api

import (
	"backend/configs"
	"backend/models"
	"backend/queries"
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
	if err != nil {
		log.Fatal(err)
	}
	
	Existance, err := queries.UserExist(db, user)
	if err != nil {
		log.Fatal(err)
	}
	if Existance {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Username already exists"))
		return
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	if err != nil {
		log.Fatal(err)
	}
	JWToken, err := configs.GenerateJWToken(&user)
	if err != nil {
		log.Fatal(err, 1)
	}
	w.Header().Set("Authorization", "Bearer "+JWToken)
	w.Header().Set("Content-Type", "application/json")

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
