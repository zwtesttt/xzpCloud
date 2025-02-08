package grpc

import (
	"google.golang.org/grpc"
)

type VmGrpcServer struct {
}

func New() *grpc.Server {
	s := grpc.NewServer()
	//vm := &VmGrpcServer{}
	//gegrpc.RegisterGreeterServer(s, vm)

	return s
}
