package middleware

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "hellogo/middleware/proto"
	"log"
	"os"
	//"os"
)

const (
	address = "127.0.0.1:21999"
)

func StartGrpcClient() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	//c := pb.NewGreeterClient(conn)
	//
	//name := "yilion"
	//if len(os.Args) > 1 {
	//	name = os.Args[1]
	//}
	//
	//r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Println(r.Message)

	c2 := pb.NewFileSaverClient(conn)

	stream, err := c2.Save(context.Background())
	if err != nil {
		fmt.Errorf("create stream failed:%v", err)
		return
	}
	defer stream.CloseSend()
	fin, err := os.ReadFile("./fd22f99815f740aba778eea680ed6c4b.jpeg")
	if err != nil {
		fmt.Errorf("send msg error:%v", err)
		return
	}

	err = stream.Send(&pb.SaveFileRequest{
		Name: "demo.jpeg",
		Data: fin,
	})
	if err != nil {
		fmt.Errorf("send msg error:%v", err)
		return
	}
}
