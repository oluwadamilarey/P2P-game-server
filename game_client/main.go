package main

import (
	"encoding/json"
	"log"
	"math"
	"math/rand"

	"github.com/gorilla/websocket"
	"github.com/oluwadamilarey/game-server/types"
)

const wsServerEndpoint = "ws://localhost:40000/ws"

type Login struct {
	ClientID int    `json:"clientID"`
	Username string `json:"username"`
}

type GameClient struct {
	conn     *websocket.Conn
	clientID int
	username string
}

func newGameClient(conn *websocket.Conn, username string) *GameClient {
	return &GameClient{
		clientID: rand.Intn(math.MaxInt),
		username: username,
		conn:     conn,
	}
}

func (c *GameClient) login() error {
	b, err := json.Marshal(types.Login{
		ClientID: c.clientID,
		Username: c.username,
	})
	if err != nil {
		return err
	}
	msg := types.WSMessage{
		Type: "login",
		Data: b,
	}
	return c.conn.WriteJSON(msg)
}

func main() {
	dialer := websocket.Dialer{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, _, err := dialer.Dial(wsServerEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	c := newGameClient(conn, "David")
	if err := c.login(); err != nil {
		log.Fatal(err)
	}
}
