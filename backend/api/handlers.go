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
		log.Println(err)
	}
	Existance, err := queries.UserExist(db, user)
	if err != nil {
		log.Println(err)
	}
	if Existance {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Username already exists"))
		return
	}
	password := user.Password
	user.Password, err = configs.HashPassword(password)
	if err != nil {
		log.Println(err)
	}

	err = queries.AddUser(db, &user)
	if err != nil {
		log.Println(err)
	}
	JWToken, err := configs.GenerateJWToken(&user)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
	}

	hashedpassword, err := queries.GetPassword(db, user.Username)
	if err != nil {
		log.Println(err)
	}
	if configs.CheckPasswordHash(user.Password, hashedpassword) {
		JWToken, err := configs.GenerateJWToken(&user)
		if err != nil {
			log.Println(err, 1)
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
		log.Println(err)
	}
	err = json.NewEncoder(w).Encode(results)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
}

func GetRooms(w http.ResponseWriter, r *http.Request) {
	db := configs.Connect_db()
	defer db.Close()

	results, err := queries.Rooms(db)
	if err != nil {
		log.Println(err)
	}
	err = json.NewEncoder(w).Encode(results)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
}


func GetMessages(w http.ResponseWriter, r *http.Request) {
	db := configs.Connect_db()
	defer db.Close()
	roomName := r.URL.Query().Get("roomName")
	messages, err := queries.Messages(db, roomName)
	if err != nil {
		log.Println(err)
	}
	username := r.Context().Value(models.UserName).(string)
	data := map[string]interface{}{
		"Messages":    messages,
		"currentUser": username,
	}
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
}
