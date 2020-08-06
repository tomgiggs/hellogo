package middleware

import (
	"github.com/funny/link"
	"github.com/funny/link/codec"

	//"github.com/funny/link/codec"
	pb "hellogo/middleware/proto"
	"log"
)


func LinkClientStart() {
	//json := codec.Json()
	//json.Register(AddReq{})
	//json.Register(AddRsp{})
	protoBuf := ProtoBuf()
	protoBuf.Register(pb.HelloRequest{})
	protoBuf.Register(pb.HelloReply{})
	clientSession, err := link.Dial("tcp","localhost:9999",  protoBuf, 1024)
	checkErr(err)
	clientSessionLoop(clientSession)
}

func clientSessionLoop(session *link.Session) {
	for i := 0; i < 10; i++ {
		err := session.Send(&pb.HelloRequest{
			Name:"jackson",
		})
		checkErr(err)
		log.Printf("Send: %d + %d", i, i)

		rsp, err := session.Receive()
		checkErr(err)
		log.Printf("Receive: %v", rsp.(*pb.HelloReply))
	}
}

func LinkJsonClientStart() {
	json := codec.Json()
	json.Register(AddReq{})
	json.Register(AddRsp{})
	clientSession, err := link.Dial("tcp","localhost:9999",  json, 1024)
	checkErr(err)
	for i := 0; i < 10; i++ {
		err := clientSession.Send(&AddReq{
		20,30,
		})
		checkErr(err)
		log.Printf("Send: %d + %d", i, i)

		rsp, err := clientSession.Receive()
		checkErr(err)
		log.Printf("Receive: %v", rsp.(*AddRsp))
	}
}