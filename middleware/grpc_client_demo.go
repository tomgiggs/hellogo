package middleware

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "hellogo/middleware/proto"
	"log"
	"os"
)

const (
	address = "127.0.0.1:18999"
)

func StartGrpcClient() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	name := "yilion"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println(r.Message)
}
