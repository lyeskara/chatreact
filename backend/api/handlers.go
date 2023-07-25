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
	password := user.Password
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

	hashedpassword, err := queries.GetPassword(db, user.Username)
	if err != nil {
		log.Println(err)
	}
	if configs.CheckPasswordHash(user.Password, hashedpassword) {
		JWToken, err := configs.GenerateJWToken(&user)
		if err != nil {
			log.Fatal(err, 1)
		}
		w.Header().Set("Authorization", "Bearer "+JWToken)
		w.Header().Set("Content-Type", "application/json")
	} else {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("password is wrong. try again"))
		return
	}

}


func GetUsers(w http.ResponseWriter, r *http.Request) {

	db := configs.Connect_db()
	defer db.Close()
	var query models.Search
	err := json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		log.Println(err)
	}
	results, err := queries.Search(db, query)
	if err != nil {
		log.Println(err, "    ", "SEARCHING")
	}
	err = json.NewEncoder(w).Encode(results)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
}

func GetRooms(w http.ResponseWriter, r *http.Request){
	db := configs.Connect_db()
	defer db.Close()

	results, err := queries.Rooms(db)
	if err != nil {
		log.Println(err, "    ", "SEARCHING")
	}
	err = json.NewEncoder(w).Encode(results)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
}

func GetMessages(w http.ResponseWriter, r *http.Request){
	db := configs.Connect_db()
	defer db.Close()
    roomName := r.URL.Query().Get("roomName")
	results, err := queries.Messages(db, roomName)
	if err != nil {
		log.Println(err, "    ", "ROOMNAME")
	}
	err = json.NewEncoder(w).Encode(results)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
}