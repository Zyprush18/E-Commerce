package services

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Zyprush18/E-Commerce/configs"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)


const (
	Success = http.StatusOK
	Created = http.StatusCreated
	NotFound = http.StatusNotFound
	Forbidden = http.StatusForbidden
	BadRequest = http.StatusBadRequest
	Unauthorized  = http.StatusUnauthorized
	MethodNotAllowed = http.StatusMethodNotAllowed
	InternalServerError = http.StatusInternalServerError
)



var jwtsecret =  []byte(configs.Config().Token)
var refresh = []byte(configs.Config().Refresh)

type Message struct {
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
	Error any `json:"error,omitempty"`
	Token string `json:"token,omitempty"`
	RefreshToken string `json:"refreshtoken,omitempty"`
}


type UserRequest struct {
	Name string `json:"name" validate:"required"` 
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginReq struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Claim struct {
	Email string
	Password string
	Role_id int
	jwt.RegisteredClaims
}


func GenerateToken(email,password string, role int) (string,string,error) {
	expireTime := time.Now().Add(15 * time.Minute)
	claim :=  &Claim{
		Email: email,
		Password: password,
		Role_id: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)

	accesstoken,err := token.SignedString(jwtsecret)
	if  err != nil {
		return "","", err
	}


	refreshexpired := time.Now().Add(7 * 24 * time.Hour)
	refreshClaims :=  &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(refreshexpired),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256,refreshClaims)
	tokenrefresh,err := refreshToken.SignedString(refresh)
	if err != nil {
		return "","",err
	}

	return accesstoken,tokenrefresh,nil
}


// func DecryptToken()  {
	
// }


func Validation(valreq interface{}) []string {
	// menginisialisasi library validator
	validate := validator.New()

	// mengecek apakah ada validasi yang error
	if err := validate.Struct(valreq); err != nil {
		var messagevalidate []string

		// perulangan yang di lakukan untuk mendapatkan error field dan tag yang kemudian di kirim ke array string
		for _, v := range err.(validator.ValidationErrors) {
			messagevalidate = append(messagevalidate, fmt.Sprintf("Field %s is %s", v.Field(), v.Tag()))
		}
		return messagevalidate
	}

	return nil
}