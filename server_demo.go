package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync/atomic"
)

var count int64

func ws(w http.ResponseWriter, r *http.Request) {
	// Upgrade connection
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	new := atomic.AddInt64(&count, 1)
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
		log.Printf("msg: %s", string(msg))
	}
}

func main() {
	// Increase resources limitations,only work under linux 
	// var rLimit syscall.Rlimit
	// if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
	// 	panic(err)
	// }
	// rLimit.Cur = rLimit.Max
	// if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
	// 	panic(err)
	// }

	// Enable pprof hooks
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
//-------------------
// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/websocket"
// )

// func ws(w http.ResponseWriter, r *http.Request) {
// 	// Upgrade connection
// 	upgrader := websocket.Upgrader{}
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		return
// 	}
// 	// Read messages from socket
// 	for {
// 		_, msg, err := conn.ReadMessage()
// 		if err != nil {
// 			conn.Close()
// 			return
// 		}
// 		log.Printf("msg: %s", string(msg))
// 	}
// }

// func mainx() {
// 	http.HandleFunc("/", ws)
// 	if err := http.ListenAndServe(":8000", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }
//-----------------------------
// import (
// 	"io"
// 	"net/http"
// )

// func hello(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "Hello GopherCon Israel 2019!")
// }

// func main() {
// 	http.HandleFunc("/", hello)
// 	http.ListenAndServe(":8000", nil)
// }
