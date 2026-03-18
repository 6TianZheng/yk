package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"lx/srv/basic/config"
	_ "lx/srv/basic/init"
	initpkg "lx/srv/basic/init"
	"lx/srv/basic/proto"
	"lx/srv/handler/service"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {

	log.Println("Consul初始化成功")
	services, err := initpkg.GetServiceWithLoadBalancer(config.GlobalConfig.Consul.ServiceName)
	if err != nil {
		log.Printf("获取用户服务失败: %v", err)
	} else {
		log.Printf("获取到用户服务: %s, 地址: %s:%d", services.Service, services.Address, services.Port)
	}
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterProductServiceServer(s, &service.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	err = initpkg.ConsulShutdown()
	if err != nil {
		return
	}
	fmt.Println("服务已退出")
}
