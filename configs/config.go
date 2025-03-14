package configs

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
)


type Configs struct {
	AppName string
	Db_User string
	Db_Pass string
	Db_Host string
	Db_Port string
	Token string
	Refresh string
}

var (
	cnf  *Configs
	// digunakan untuk memastikan bahwa sebuah fungsi hanya dijalankan sekali selama aplikasi berjalan, bahkan jika dipanggil berkali-kali dari berbagai tempat.
	once sync.Once 
)

func Config() *Configs {
	rootPath, err := filepath.Abs(".") // Dari "configs/" naik satu level ke root proyek
	if err != nil {
		log.Fatal("Error getting root path:", err)
	}

	// Load .env dari root proyek
	envPath := filepath.Join(rootPath, ".env")

	once.Do(func ()  {
		if err := godotenv.Load(envPath); err != nil {
			log.Fatal(err.Error())
			log.Fatal("Error loading .env file")
		}
	
		cnf = &Configs{
			AppName: os.Getenv("APP_NAME"),
			Db_User: os.Getenv("USER_DATABASE"),
			Db_Pass: os.Getenv("PASS_DATABASE"),
			Db_Host: os.Getenv("HOST"),
			Db_Port: os.Getenv("PORT"),
			Token: os.Getenv("JWT_SECRET"),
			Refresh: os.Getenv("REFRESH_SECRET"),
		}
	})

	return cnf
	
}