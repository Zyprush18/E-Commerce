package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Zyprush18/E-Commerce/common"
	// "github.com/Zyprush18/E-Commerce/configs"
	"github.com/Zyprush18/E-Commerce/services"
	pb "github.com/Zyprush18/E-Commerce/services/user-service/proto"
)

// menginisialisasi pointer dari struct yang ada di folder common
var grpcClient *common.GRPCCLIENT

func InitGRPCCLIENT() {
	grpcClient = common.NewGRPCCLIENT()
}

func Register(w http.ResponseWriter, r *http.Request) {
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

	// mengecek validasi
	if err := services.Validation(userReq); err != nil {
		w.WriteHeader(services.BadRequest)
		json.NewEncoder(w).Encode(services.Message{
			Message: "Validation Error",
			Error:   err,
		})
		return
	}

	// membuat context kosong yang sering di gunakan untuk pemanggilan fungsi gRPC
	ctx := context.Background()

	// memanggil metode register yang ada di UserService pada gRPC Server dan mengirimkan request yang berisi data pengguna. ctx (context) di gunakan untuk mengatur lifecycle request
	userClient, err := grpcClient.RegisterService.Register(ctx, &pb.ReqRegister{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: userReq.Password,
	})

	// mengirimkan response jika terjadi error pada metode register
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(services.Message{
			Message: "Email yang Anda masukkan sudah digunakan. Harap gunakan email lain.",
		})
		return
	}

	// mengembalikan message dari metode register yang ada di gRPC server
	w.WriteHeader(services.Success)
	json.NewEncoder(w).Encode(services.Message{
		Message: userClient.Message,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(services.MethodNotAllowed)
		json.NewEncoder(w).Encode(services.Message{
			Message: "Method Not Allowed",
		})
		return
	}

	loginReq := services.LoginReq{}

	// mengecek apakah body nya sudah sama dengan user request
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		w.WriteHeader(services.BadRequest)
		json.NewEncoder(w).Encode(services.Message{
			Message: "Invalid Field",
		})
		return
	}

	if err := services.Validation(loginReq); err != nil {
		w.WriteHeader(services.BadRequest)
		json.NewEncoder(w).Encode(services.Message{
			Message: "Validation Error",
			Error:   err,
		})
		return
	}

	ctx := context.Background()
	login, err := grpcClient.LoginService.Login(ctx, &pb.ReqLogin{
		Email:    loginReq.Email,
		Password: loginReq.Password,
	})

	if err != nil {
		w.WriteHeader(services.BadRequest)
		json.NewEncoder(w).Encode(services.Message{
			Message: "Email Atau Password Salah",
		})
		return
	}

	
	w.WriteHeader(services.Success)
	json.NewEncoder(w).Encode(services.Message{
		Message: login.Message,
		Data: login.Data,
		Token:        login.Token,
		RefreshToken: login.Refresh,
	})
}

func Logout(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(services.MethodNotAllowed)
		json.NewEncoder(w).Encode(services.Message{
			Message: "Method not allowed",
		})
		return
	}

		// Ambil path contoh: /logout/123
		path := r.URL.Path
		parts := strings.Split(path, "/")
	
		// Pastikan path memiliki ID di urutan ketiga
		if len(parts) < 3 || parts[2] == "" {
			log.Println("tidak ada id")
			return
		}
	
		userid := parts[2] // Ambil {id} dari URL
		

	ctx := context.Background()
	logouts,err := grpcClient.LogoutService.Logout(ctx, &pb.LogoutRequest{
		Id: userid,
	})

	if err != nil {
		w.WriteHeader(services.BadRequest)
		json.NewEncoder(w).Encode(services.Message{
			Message: "User Tidak Di Temukan",
			Error: err.Error(),
		})
		return
	}

	w.WriteHeader(services.Success)
	json.NewEncoder(w).Encode(services.Message{
		Message: logouts.Message,
	})

}


