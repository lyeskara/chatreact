package models


type CtxKey string


const UserName CtxKey = "username"

type Search struct {
	Search string `json:"search"`
}

type Searched struct {
	Username string `json:"username"`
}
