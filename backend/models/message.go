package models

type Message struct{
	Message string `json:"message"`
	Room string `json:"roomName"`
	Username string `json:"username"`
}