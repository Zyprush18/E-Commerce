package main

import (
	"log"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("User Service Is Running"))
}


func main()  {
	http.HandleFunc("/user", GetUser)


	log.Println("Api Gateway Running On Port 8080")
	http.ListenAndServe(":8080",nil)
}