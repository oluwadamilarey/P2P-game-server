package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"

	"github.com/anthdm/hollywood/actor"
	"github.com/gorilla/websocket"
	"github.com/oluwadamilarey/game-server/types"
)

type HTTPServer struct{}

type PlayerSession struct {
	sessionID int
	clientID  int
	inLobby   bool
	conn      *websocket.Conn
}

type GameServer struct {
	ctx      *actor.Context
	sessions map[*actor.PID]struct{}
}

func newPlayerSession(sid int, conn *websocket.Conn) actor.Producer {
	return func() actor.Receiver {
		return &PlayerSession{
			sessionID: sid,
			conn:      conn,
			inLobby:   true,
		}
	}
}

func (s *PlayerSession) readLoop() {
	var msg types.WSMessage
	for {
		if err := s.conn.ReadJSON(msg); err != nil {
			fmt.Println()
			return
		}
		go s.handleMessage(msg)
	}
}

func (s *PlayerSession) handleMessage(msg types.WSMessage) {
	switch msg.Type {
	case "login":
		var loginMsg types.Login
		if err := json.Unmarshal(msg.Data, &loginMsg); err != nil {
			panic(err)
		}
		fmt.Println(loginMsg)
	}
}

// Implement the Receive method for PlayerSession to satisfy the actor.Receiver interface.
func (s *PlayerSession) Receive(c *actor.Context) {
	// Handle messages received by PlayerSession here.
	switch c.Message().(type) {
	case actor.Started:
		s.readLoop()
	}
}

func newGameServer() actor.Receiver {
	return &GameServer{
		sessions: make(map[*actor.PID]struct{}),
	}
}

func (s *GameServer) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		s.startHTTP()
		s.ctx = c
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
	sid := rand.Intn(math.MaxInt)
	// for each client thats going to make a connection with the server, to spawn a new child server
	pid := s.ctx.SpawnChild(newPlayerSession(sid, conn), fmt.Sprintf("session_%d", sid))
	fmt.Printf("client with sid %d and pid %s just connected\n", sid, pid)
}

func main() {
	e := actor.NewEngine()
	e.Spawn(newGameServer, "server")
	select {}
}
