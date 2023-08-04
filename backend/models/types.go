package models

import "os"

type CtxKey string

var Secret string = os.Getenv("TOKEN_KEY")

const UserName CtxKey = "username"

type Search struct {
	Search string `json:"search"`
}

type Searched struct {
	Username string `json:"username"`
}
