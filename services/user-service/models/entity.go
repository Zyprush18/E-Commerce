package models


type UserModel struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role_id uint  `json:"role_id"`
}

type Roles struct {
	Id uint `json:"id"`
	Role string `json:"role"`
	User []UserModel
}