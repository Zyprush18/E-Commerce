package repository

import (
	"fmt"
	"log"

	"github.com/Zyprush18/E-Commerce/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	config := configs.Config()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/user_db?charset=utf8mb4&parseTime=True&loc=Local",
		config.Db_User, config.Db_Pass, config.Db_Host, config.Db_Port)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Simpan koneksi DB global
	DB = db

	// Migrasi tabel sekaligus
	if err := DB.AutoMigrate(&Roles{}, &User{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Database connected and migration successful!")

}
