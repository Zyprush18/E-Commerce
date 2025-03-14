package routes

import (
	"log"
	"net/http"

	"github.com/Zyprush18/E-Commerce/api-gateway/handlers/auth"

)

func Routes()  {
	auth.InitGRPCCLIENT()

	http.HandleFunc("/register", auth.Register)
	http.HandleFunc("/login",auth.Login)


	log.Println("Api gateway running on port : 8080")
	http.ListenAndServe(":8080",nil)
}