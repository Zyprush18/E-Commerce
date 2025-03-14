package handler

import (
	"context"
	// "strings"

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

func (s *Login) Login(ctx context.Context, req *pb.ReqLogin) (*pb.ResLogin, error)  {
	user := &models.UserModel{}

	if err := repository.DB.Table("users").Where("email=?",req.Email).Find(&user).Error;err != nil {
		return nil,err
	}

	if err := service.DecryptPassword(user.Password,req.Password);err != nil {
		return nil, err
	}

	accesstoken,refreshtoken,err := services.GenerateToken(req.Email,req.Password,int(user.Role_id))
	if err != nil {
		return nil, err
	}
	
	return &pb.ResLogin{
		Message: "Berhasil Login",
		Data: []string{user.Name,user.Email},
		Token: accesstoken,
		Refresh: refreshtoken,
	},nil



}