package middleware

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	pb "hellogo/middleware/proto"

	"net"
	"time"
)
type server struct {
	pb.UnimplementedGreeterServer
}

func StartGrpcServer() {
	keepParams := grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     time.Duration(time.Second * 60),
		MaxConnectionAgeGrace: time.Duration(time.Second * 20),
		Time:                  time.Duration(time.Second * 60),
		Timeout:               time.Duration(time.Second * 20),
		MaxConnectionAge:      time.Duration(time.Hour * 2),
	})
	srv := grpc.NewServer(keepParams)
	lis, err := net.Listen("tcp", "127.0.0.1:8999")
	pb.RegisterGreeterServer(srv,&server{})
	if err != nil {
		fmt.Println("start grpc server error: ",err)
		panic(err)
	}
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
	//go func() {
	//	if err := srv.Serve(lis); err != nil {
	//		panic(err)
	//	}
	//}()
}

// protoc --go-grpc_out=proto    -I proto/ proto/hellogrpc.proto
