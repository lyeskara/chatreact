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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool) // Connected clients
var broadcast = make(chan []byte)            // Channel to broadcast messages to clients

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	db := configs.Connect_db()
	// Upgrade the HTTP connection to a WebSocket connection.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	clients[conn] = true

	for {
		// Read message from the client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			delete(clients, conn)
			break // Exit the loop to close the connection
		}
		// Parse the message json format
		var data map[string]interface{}
		err = json.Unmarshal(msg, &data)
		if err != nil {
			log.Println("failed to marshal the json message from websocket")
		}
		Category, ok := data["type"].(string)
		if !ok {
			log.Println("invalid data for type of message sent through websocket")
		}
		// process data based on request type
		switch Category {

		case "newRoom":
			var room models.Room
			room.RoomName = data["content"].(string)
			err = queries.AddRoom(db, &room)
			if err != nil {
				log.Println(err)
			}
			
			broadcast <- []byte(room.RoomName)

		case "newMessage":
			var message models.Message
			message.Username = r.Context().Value(models.UserName).(string)
			message.Message = data["msg"].(string)
			message.Room = data["room"].(string)
			err = queries.AddMessage(db, &message)
			if err != nil {
				log.Println(err)
			}
			MessageSent, err := json.Marshal(message)
			if err != nil {
				log.Println(err)
			}
			// Write message back to the client
			broadcast <- MessageSent
		}
	}
}

func HandleMessages() {
	for {
		message := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("Error writing message to WebSocket:", err)
				handleClientError(client)
			}
		}
	}
}

// Function to handle client errors and disconnect
func handleClientError(client *websocket.Conn) {
	client.Close()
	delete(clients, client)
}
