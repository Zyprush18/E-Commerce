package main

import (
	"log"
	"net/http"

	"github.com/Zyprush18/E-Commerce/services/user-service/repository"
)

func GetUsers(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("User Service Is Running"))
}

func main() {
	// migrateion
	repository.Connect()


	http.HandleFunc("/users",GetUsers)

	log.Println("User Service Running On Port 8081")
	http.ListenAndServe(":8081",nil)
}