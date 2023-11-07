package main

import "github.com/anthdm/hollywood/actor"

type GameServer struct{}

func newGameServer() actor.Receiver {
    return &GameServer{}
}

func (s *GameServer) Recieve(c *actor.Context)  {}

func  main() {
    e := actor.NewEngine()
    e.Spawn(newGameServer, "server")
}