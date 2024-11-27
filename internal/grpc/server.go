package grpc

import (
	main_v1 "github.com/pelicanch1k/grpcProduct-proto/proto/gen/go/proto/main"
	"google.golang.org/grpc"
)

type serverApi struct {
	main_v1.UnimplementedProductServiceServer
}

func Register(gRPC *grpc.Server) {
	main_v1.RegisterProductServiceServer(gRPC, &serverApi{})
}
