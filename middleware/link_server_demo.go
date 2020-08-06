package middleware

import (
	"fmt"
	"github.com/funny/link"
	"github.com/funny/link/codec"
	pb "hellogo/middleware/proto"
	"log"
	"net"
)

type AddReq struct {
	A, B int
}

type AddRsp struct {
	C int
}

type Server struct{}

func (*Server) HandleSession(session *link.Session) {
	for {
		req, err := session.Receive()
		fmt.Println(req)
		checkErr(err)
		err = session.Send(&pb.HelloReply{
			Message: " this is server",
		})
		checkErr(err)
	}
}
func LinkJsonStart() {
	json := codec.Json()
	json.Register(AddReq{})
	json.Register(AddRsp{})
	listen, err := net.Listen("tcp", "127.0.0.1:9999")
	checkErr(err)
	server := link.NewServer(listen, json, 1024, new(Server))
	server.Serve()
}

func LinkStart() {
	protoBuf := ProtoBuf()
	protoBuf.Register(pb.HelloRequest{})
	protoBuf.Register(pb.HelloReply{})

	listen, err := net.Listen("tcp", "127.0.0.1:9999")
	checkErr(err)
	server := link.NewServer(listen, protoBuf, 1024, new(Server))
	server.Serve()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
