package main

import (
	"fmt"
	"net/http"

	"github.com/anthdm/hollywood/actor"
	"github.com/gorilla/websocket"
)

type HTTPServer struct{}

type PlayerSession struct {
	clientID int
	username string
	inLobby  bool
	conn     *websocket.Conn
}

type GameServer struct {
	ctx *actor.Context
}

func newPlayerSession(clientID int, username string, conn *websocket.Conn) actor.Producer {
	return func() actor.Receiver {
		return &PlayerSession{
			clientID: clientID,
			username: username,
			conn:     conn,
			inLobby:  true,
		}
	}
}

// Implement the Receive method for PlayerSession to satisfy the actor.Receiver interface.
func (ps *PlayerSession) Receive(c *actor.Context) {
	// Handle messages received by PlayerSession here.
}

func newGameServer() actor.Receiver {
	return &GameServer{}
}

func (s *GameServer) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		s.startHTTP()
		_ = msg
	}
}

func (s *GameServer) startHTTP() {
	fmt.Println("starting HTTP server on port -> 40000")
	go func() {
		http.HandleFunc("/ws", s.handleWS)
		http.ListenAndServe(":40000", nil)
	}()
}

//upgrades the HTTP server connection to the WebSocket protocol.
func (s *GameServer) handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		fmt.Println("ws upgrade err:", err)
		return
	}
	fmt.Print("Novel Client trying to connect")
	fmt.Print(conn)
}

func main() {
	e := actor.NewEngine()
	e.Spawn(newGameServer, "server")
	select {}
}
