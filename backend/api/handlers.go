package api

import (
	"backend/configs"
	"backend/models"
	"backend/queries"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
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

func Userdata(w http.ResponseWriter, r *http.Request) {
	db := configs.Connect_db()
	defer db.Close()

	username := r.Context().Value(models.UserName).(string)
	var user models.User
	user.Username = username

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


func HandleWebSocket(w http.ResponseWriter, r *http.Request) {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	// Upgrade the HTTP connection to a WebSocket connection.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		// Read message from the client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		
		// Print the received message
		log.Printf("Received message: %s\n", msg)

		// Write message back to the client
		err = conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println(err)
			break
		}
	}
}

