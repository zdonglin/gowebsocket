package SocketService

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var Conn *websocket.Conn

type Data struct {
	Session     string
	Sessiontype int
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	for {
		mt, session, err := c.ReadMessage()

		fmt.Println("session:", string(session))

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("mt:", mt)

		if err != nil {
			log.Println("read:", err)
			break
		}
		SetConn(string(session), c)
		SetType(string(session), mt)
	}
}

func Init(url string) {
	go func() {
		var addr = flag.String("addr", url, "http service address")
		http.HandleFunc("/echo", echo)
		log.Fatal(http.ListenAndServe(*addr, nil))
	}()
}
