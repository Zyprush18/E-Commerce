package handler

import (
	"context"

	pb "github.com/Zyprush18/E-Commerce/services/user-service/proto"
	"github.com/Zyprush18/E-Commerce/services/user-service/models"
	"github.com/Zyprush18/E-Commerce/services/user-service/repository"
	"github.com/Zyprush18/E-Commerce/services/user-service/service"

)

type Register struct {
	// struct bawaan yang dihasilkan oleh protobuf saat melakukan generate kode gRPC
	pb.UnimplementedRegisterServiceServer
}


func (s *Register) Register(ctx context.Context, req *pb.ReqRegister) (*pb.ResRegister,error) {
	// hashing password
	hashingpw, err := service.HashingPassword(req.Password)
	if err != nil{
		return nil, err
	}

	register := &models.UserModel{
		Name: req.Name,
		Email: req.Email,
		Password: hashingpw,
		Role_id: 2,
	}

	if err := repository.DB.Table("users").Create(register).Error;err != nil {
		return nil, err
	}

	return &pb.ResRegister{
		Message: "Berhasil Register",
	},nil

}