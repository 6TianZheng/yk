package initpkg

import (
	"flag"
	"log"
	initpkg "lx/srv/basic/init"

	"lx/bff/basic/config"
	__ "lx/srv/basic/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	GrpcInit()
	initpkg.MysqlInit()
}

func GrpcInit() {
	flag.Parse()
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	config.ProductClient = __.NewProductServiceClient(conn)
}
