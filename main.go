package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

const wsServerEndpoint = "ws://localhost:4000/ws"

func main() {
    dialer := websocket.Dialer{
        ReadBufferSize: 1024,
        WriteBufferSize: 1024,
    }
    conn, _,err  := dialer.Dial(wsServerendpoint)
    if err != nil {
        log.Fatal(err)
    }


    fmt.Println("vim-go")
}