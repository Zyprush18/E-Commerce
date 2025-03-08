package repository



type User struct {
	Id uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role_id uint `json:"id_role"`
}

type Roles struct {
	Id uint `json:"id" gorm:"primaryKey"`
	Role string `json:"role"`
	User []User `gorm:"foreignKey:Role_id;references:id"`
}