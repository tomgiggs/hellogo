package middleware
/*
从这个项目复制过来的：https://github.com/eranyanay/1m-go-websockets
主要用于学习，感谢作者
 */

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync/atomic"
)

var count int64

//每次有新的连接都会调用一次这个函数
func ws(w http.ResponseWriter, r *http.Request) {
	// Upgrade connection
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	new := atomic.AddInt64(&count, 1)//原子操作包，用于并发环境下的锁操作
	if new%100 == 0 {
		log.Printf("Total number of connections: %v", new)
	}
	defer func() {
		new := atomic.AddInt64(&count, -1)
		if new%100 == 0 {
			log.Printf("Total number of connections: %v", new)
		}
	}()

	// Read messages from socket
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		conn.WriteJSON("goooooooooooood")
		log.Printf("msg: %s", string(msg))
	}
}

func WebSocketDemo() {

	go func() {
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			log.Fatalf("Pprof failed: %v", err)
		}
	}()
    //为什么监听了两个端口呢
	http.HandleFunc("/", ws)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}


