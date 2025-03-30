package handler

import (
	"context"

	"github.com/Zyprush18/E-Commerce/configs"
	"github.com/Zyprush18/E-Commerce/services"
	"github.com/Zyprush18/E-Commerce/services/user-service/models"
	pb "github.com/Zyprush18/E-Commerce/services/user-service/proto"
	"github.com/Zyprush18/E-Commerce/services/user-service/repository"
	"github.com/Zyprush18/E-Commerce/services/user-service/service"
)

type Register struct {
	// struct bawaan yang dihasilkan oleh protobuf saat melakukan generate kode gRPC
	pb.UnimplementedRegisterServiceServer
}

type Login struct {
	pb.UnimplementedLoginServiceServer
}

type Logout struct {
	pb.UnimplementedLogoutServiceServer
}

func (s *Register) Register(ctx context.Context, req *pb.ReqRegister) (*pb.ResRegister, error) {
	// hashing password
	hashingpw, err := service.HashingPassword(req.Password)
	if err != nil {
		return nil, err
	}

	register := &models.Register{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashingpw,
		Role_id:  2,
	}

	if err := repository.DB.Table("users").Create(register).Error; err != nil {
		return nil, err
	}

	return &pb.ResRegister{
		Message: "Berhasil Register",
	}, nil
}

func (s *Login) Login(ctx context.Context, req *pb.ReqLogin) (*pb.ResLogin, error) {
	user := &models.UserModel{}

	if err := repository.DB.Table("users").Where("email= ?", req.Email).Find(&user).Error; err != nil {
		return nil, err
	}

	if err := service.DecryptPassword(user.Password, req.Password); err != nil {
		return nil, err
	}

	accesstoken, refreshtoken, err := services.GenerateToken(req.Email, req.Password, int(user.Role_id))
	if err != nil {
		return nil, err
	}

	configs.KeepToRedis(user.Id, accesstoken,refreshtoken)
	
	userdata := map[string]string{
		"id":    user.Id,
		"name":  user.Name,
		"email": user.Email,
	}


	return &pb.ResLogin{
		Message: "Berhasil Login",
		Data: userdata,
		Token:   accesstoken,
		Refresh: refreshtoken,
	}, nil

}

func (s *Logout) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error)  {
	user_id := req.GetId()

	pesan, err := configs.Logout(user_id)
	if err != nil {
		return nil, err
	}

	return &pb.LogoutResponse{
		Message: pesan,
	},nil
}
