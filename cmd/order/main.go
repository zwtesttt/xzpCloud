package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	ordergrpc "github.com/zwtesttt/xzpCloud/internal/order/api/grpc"
	"github.com/zwtesttt/xzpCloud/internal/order/api/handler"
)

var (
	shutdown = make(chan struct{}, 1)
)

func main() {
	var (
		r       = handler.New()
		grpcSvc = ordergrpc.New()
		httpSvc = &http.Server{
			Addr:    ":8080",
			Handler: r,
		}
	)

	go startHttp(r)
	go startGrpc(grpcSvc)
	go gracefullyShutdown(context.Background(), grpcSvc, httpSvc)

	<-shutdown
}

func gracefullyShutdown(ctx context.Context, grpcSvc *grpc.Server, r *http.Server) {
	// 创建一个通道，用于接收信号
	sigChan := make(chan os.Signal, 1)
	// 注册要监听的信号,在控制台中输入 ctrl+c 可以发送信号其他信号也是差不多
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	//做一些事情
	grpcSvc.Stop()

	//停止http服务器
	err := r.Shutdown(ctx)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	close(shutdown)
}

func startGrpc(grpcSvc *grpc.Server) {
	//grpc服务器
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("grpc server start")
	err = grpcSvc.Serve(listen)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

func startHttp(r *handler.Handler) {
	//http服务器
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
