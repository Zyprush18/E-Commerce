package repository

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


func Connect()  {
	var err error
	const mysql_driver = "root:root@tcp(127.0.0.1:3306)/user_db?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := mysql_driver


	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Connection Failed")
	}
	DB.AutoMigrate(&Roles{})	
	DB.AutoMigrate(&User{})	

	log.Println("Migration Success")

}


