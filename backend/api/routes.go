package api


import (
	"github.com/go-chi/chi/v5"	
)


func Routes(r *chi.Mux){

	r.Post("/signup", Signup)
	r.Post("/login", Login)
}