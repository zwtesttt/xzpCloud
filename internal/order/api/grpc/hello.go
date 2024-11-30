package grpc

import (
	"context"
	"fmt"
	gegrpc "github.com/zwtesttt/xzpCloud/internal/order/api/grpc/pb"
)

func (o *OrderGrpcServer) SayHello(ctx context.Context, request *gegrpc.HelloRequest) (*gegrpc.HelloReply, error) {
	fmt.Println("调用grpc方法")
	return &gegrpc.HelloReply{Message: "hello " + request.Name}, nil
}
