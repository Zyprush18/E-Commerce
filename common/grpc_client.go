package common

import (
	"log"

	pb "github.com/Zyprush18/E-Commerce/services/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// menginisialisasi Register Service untuk client
type GRPCCLIENT struct {
	RegisterService pb.RegisterServiceClient
	LoginService pb.LoginServiceClient

}

func NewGRPCCLIENT() *GRPCCLIENT {
	// membuat koneksi ke gRPC server yang berjalan di port 8081 dan untuk grpc.WithTransportCredentials(insecure.NewCredentials()) untuk pengujian di lokal (jika production maka kita harus menganti grpc.WithTransportCredentials(insecure.NewCredentials()) dengan sertifikat TLS/SSL)
	conn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))

	// mengecek apabila gagal terhubung dengan gRPC client
	if err != nil {
		log.Fatalf("Failed to Connect to gRPC server: %v", err)
	}

	// mengembalikan instance dari gRPC Client yang dapat digunakan untuk berkomunikasi dengan gRPC Server.
	return &GRPCCLIENT{
		RegisterService: pb.NewRegisterServiceClient(conn),
		LoginService: pb.NewLoginServiceClient(conn),
	}
}
