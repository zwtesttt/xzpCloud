package grpc

import (
	gegrpc "github.com/zwtesttt/xzpCloud/internal/order/api/grpc/pb"
	"google.golang.org/grpc"
)

var _ gegrpc.GreeterServer = &OrderGrpcServer{}

type OrderGrpcServer struct {
	gegrpc.UnimplementedGreeterServer
}

func New() *grpc.Server {
	s := grpc.NewServer()
	vm := &OrderGrpcServer{}
	gegrpc.RegisterGreeterServer(s, vm)

	return s
}
