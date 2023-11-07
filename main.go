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
    conn    *websocket.Conn
    clientID int
    username string
}

func newGameClient(conn *websocket.Conn, username string) *GameClient {
    return &GameClient{
        clientID: rand.Intn(Math.MaxInt),
        username: username,
    }
}

func (c *GameClient) login() error {
    return c.conn.WriteJSON(Login{
        ClientID: c.clientID,
        Username: c.username,
    })
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
    c := newGameClient(conn, "David")
   if err := c.Login(); err != nil {

   }
    for(

    )
    fmt.Println("vim-go")
}

