package middleware

import (
	"context"
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

func (*server) SayHello(context.Context, *pb.HelloRequest) (*pb.HelloReply, error){
	return &pb.HelloReply{
		Message:" this is server",
	}, nil
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
	// 在gRPC中有两种拦截器UnaryInterceptor和StreamInterceptor，其中UnaryInterceptor拦截普通的一次请求一次响应的rpc服务，StreamInterceptor拦截流式的rpc服务。
	//srv := grpc.NewServer(grpc.UnaryInterceptor(LogicIntInterceptor)) //这里可以注册拦截器，对请求进行授权认证等操作
	lis, err := net.Listen("tcp", "127.0.0.1:18999")
	pb.RegisterGreeterServer(srv,&server{})//这里传入的server对象即是在接收到请求是具体处理的函数，需要实现proto中定义的所有函数
	if err != nil {
		fmt.Println("start grpc server error: ",err)
		panic(err)
	}
	//proto.Marshal()
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
