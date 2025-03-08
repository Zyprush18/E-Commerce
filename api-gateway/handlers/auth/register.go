package auth

import (
	"context"
	"encoding/json"
	"fmt"
	// "log"
	"net/http"

	"github.com/Zyprush18/E-Commerce/common"
	"github.com/Zyprush18/E-Commerce/services"
	pb "github.com/Zyprush18/E-Commerce/services/user-service/proto"
	"github.com/go-playground/validator/v10"
)

// menginisialisasi pointer dari struct yang ada di folder common
var grpcClient *common.GRPCCLIENT

func InitGRPCCLIENT()  {
	grpcClient = common.NewGRPCCLIENT()
}

func Register(w http.ResponseWriter, r *http.Request)  {
	// set header menjadi application/json
	w.Header().Set("Content-Type", "application/json")

	// mengecek method 
	if r.Method != http.MethodPost {
		w.WriteHeader(services.MethodNotAllowed)
		json.NewEncoder(w).Encode(services.Message{
			Message: "Method Not Allowed",
		})
		return
	}

	// mengisialisasi struct user request
	userReq := services.UserRequest{}

	// mengecek apakah body nya sudah sama dengan user request 
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		w.WriteHeader(services.BadRequest)
		json.NewEncoder(w).Encode(services.Message{
			Message: "Invalid Field",
		})
		return
	}

	// menginisialisasi library validator
	validate := validator.New()

	// mengecek apakah ada validasi yang error
	if err := validate.Struct(userReq); err != nil {
		var messagevalidate []string

		// perulangan yang di lakukan untuk mendapatkan error field dan tag yang kemudian di kirim ke array string
		for _, v := range err.(validator.ValidationErrors) {
			messagevalidate = append(messagevalidate, fmt.Sprintf("Field %s is %s", v.Field(), v.Tag()))
		}

		w.WriteHeader(services.BadRequest)
		json.NewEncoder(w).Encode(services.Message{
			Message: "Validation Error",
			Error: messagevalidate,
		})
		return
	}

	// membuat context kosong yang sering di gunakan untuk pemanggilan fungsi gRPC 
	ctx := context.Background() 

	// memanggil metode register yang ada di UserService pada gRPC Server dan mengirimkan request yang berisi data pengguna. ctx (context) di gunakan untuk mengatur lifecycle request
	userClient, err := grpcClient.UserService.Register(ctx, &pb.ReqRegister{
		Name: userReq.Name,
		Email: userReq.Email,
		Password: userReq.Password,
	})

	// mengirimkan response jika terjadi error pada metode register
	if err != nil {
		// log.Fatal(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(services.Message{
			Message: "Failed Register User",
			Error: err.Error(),
		})
		return
	}

	// mengembalikan message dari metode register yang ada di gRPC server
	w.WriteHeader(services.Success)
	json.NewEncoder(w).Encode(services.Message{
		Message: userClient.Message,
	})



}