package basic_grammar

import (
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
)
func StartHttpServer(){
	SimpleHttpServer()
}
//websocket 处理器
func wsHandler(w http.ResponseWriter, r *http.Request) {

	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			return
		}
		log.Printf("msg: %s", string(msg))
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello this is simple http server")
}

func SimpleHttpServer(){
	http.HandleFunc("/", hello)
	http.HandleFunc("/ws",wsHandler)
	err:=http.ListenAndServe(":8000", nil)
	if err!=nil{
		log.Fatal("fatal error happen when start server")
		return
	}
}


