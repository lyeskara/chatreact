package api

import (
	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux) {

	r.Post("/signup", Signup)
	r.Post("/login", Login)
	r.Post("/getUsers", GetUsers)
	r.Get("/ws", HandleWebSocket)
	go HandleMessages()
	r.Get("/getRooms", GetRooms)
	r.Get("/getMessages", GetMessages)

}
