package main

import (
	"log"
	"net"

	"github.com/Zyprush18/E-Commerce/services/user-service/handler"
	pb "github.com/Zyprush18/E-Commerce/services/user-service/proto"
	"github.com/Zyprush18/E-Commerce/services/user-service/repository"
	"google.golang.org/grpc"
)

func main() {
	// Migration database
	repository.Connect()

	register := &handler.Register{}
	login := &handler.Login{}

	// Buat listener untuk gRPC
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Jalankan server gRPC
	s := grpc.NewServer()
	pb.RegisterRegisterServiceServer(s, register)
	pb.RegisterLoginServiceServer(s, login)

	log.Println("User Service Running On Port 8081")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
