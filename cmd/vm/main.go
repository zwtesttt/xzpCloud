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

	vmgrpc "github.com/zwtesttt/xzpCloud/internal/vm/api/grpc"
	"github.com/zwtesttt/xzpCloud/internal/vm/api/handler"
	"github.com/zwtesttt/xzpCloud/pkg/config"
	"github.com/zwtesttt/xzpCloud/pkg/db"
	"github.com/zwtesttt/xzpCloud/pkg/vmi"
)

var (
	shutdown = make(chan struct{}, 1)
	vmicli   vmi.VirtualMachineInterface
)

func main() {
	configPath := os.Getenv("VM_CONFIG")
	if configPath == "" {
		configPath = "./config/vm.yaml"
	}
	cfg := config.Init(configPath)
	err := initClient(cfg)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var (
		r       = handler.New(vmicli)
		grpcSvc = vmgrpc.New(cfg)
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

	fmt.Println("vm grpc server start on :8081")
	err = grpcSvc.Serve(listen)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

func startHttp(r *handler.Handler) {
	//http服务器
	fmt.Println("vm http server start on :8080")
	err := r.Run(":8080")
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

	fmt.Println("cfg", cfg.KubeConfig)
	vmicli = vmi.NewVirtHandler(cfg)
	return nil
}
