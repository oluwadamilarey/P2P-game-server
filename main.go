package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Login struct {
    ClientID int `json:"clientID"`
    Username string `json:"username"`
}


type GameClient struct {
    clientID int
    username string
}

func NewGameClient(username string) *GameClient {
    return &GameClient{
        clientID: rand.Intn(Math.MaxInt),
        username: username,
    }
}

const wsServerEndpoint = "ws://localhost:4000/ws"

func main() {
    dialer := websocket.Dialer{
        ReadBufferSize: 1024,
        WriteBufferSize: 1024,
    }
    conn, _,err  := dialer.Dial(wsServerEndpoint, nil)
    if err != nil {
        log.Fatal(err)
    }

    for(

    )
    fmt.Println("vim-go")
}

func login(conn *websocket.Conn, data Login) error {
    return conn.WriteJSON(data)
}