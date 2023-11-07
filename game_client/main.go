package main

import (
	"log"
	"math"
	"math/rand"

	"github.com/gorilla/websocket"
)

const wsServerEndpoint = "ws://localhost:4000/ws"

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
        clientID: rand.Intn(math.MaxInt),
        username: username,
    }
}

func (c *GameClient) login() error {
    return c.conn.WriteJSON(Login{
        ClientID: c.clientID,
        Username: c.username,
    })
}

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
    if err := c.login(); err != nil {
        log.Fatal(err)
    }
}

