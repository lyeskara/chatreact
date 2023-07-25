package models

type Message struct{
	Message string `json:"message"`
	RoomName string `json:"roomName"`
	Username string `json:"username"`
}