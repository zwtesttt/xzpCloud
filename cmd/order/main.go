package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/zwtesttt/xzpCloud/pkg/config"
	"github.com/zwtesttt/xzpCloud/pkg/db"

	"google.golang.org/grpc"

	ordergrpc "github.com/zwtesttt/xzpCloud/internal/order/api/grpc"
	"github.com/zwtesttt/xzpCloud/internal/order/api/handler"
)

var (
	shutdown = make(chan struct{}, 1)
)

func main() {
	cfg := config.Init("./config/order.yaml")
	err := initClient(cfg)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var (
		r       = handler.New()
		grpcSvc = ordergrpc.New()
		httpSvc = &http.Server{
			Addr:    ":8086",
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
	listen, err := net.Listen("tcp", ":8087")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("order grpc server start on :8087")
	err = grpcSvc.Serve(listen)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

func startHttp(r *handler.Handler) {
	//http服务器
	fmt.Println("order http server start on :8086")
	err := r.Run(":8086")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

func initClient(cfg *config.Config) error {
	err := db.InitDatabase(cfg)
	if err != nil {
		return err
	}

	return nil
}
