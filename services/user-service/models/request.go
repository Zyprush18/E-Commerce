package models


type Register struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role_id uint  `json:"role_id"`
}

type Login struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role_id uint  `json:"role_id"`
}