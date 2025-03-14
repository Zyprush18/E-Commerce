package middleware

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/Zyprush18/E-Commerce/services"
// )

// type Logging struct {
// 	next http.Handler
// }

// func (nx *Logging) LoginMiddleware(w http.ResponseWriter, r *http.Request) {
// 	token := r.Header.Get("Authorization")

// 	if token != "" {
// 		w.Header().Add("Content-type","application/json")
// 		w.WriteHeader(services.Unauthorized)
// 		json.NewEncoder(w).Encode(services.Message{
// 			Message: "Unauthorized: Missing Token",
// 		})
// 	}



// 	nx.next.ServeHTTP(w,r)
// }