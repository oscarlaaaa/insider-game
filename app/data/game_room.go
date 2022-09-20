package data

import (
	"math/rand"
)

type GameRoom struct {
	id      int
	name    string
	players []Player
}

func createGameRoom(name string) *GameRoom {
	gameRoom := GameRoom{
		id:      rand.Intn(1000000),
		name:    name,
		players: make([]Player, 0, 4),
	}
	return &gameRoom
}

func addPlayer(player *Player, gameRoom *GameRoom) {
	gameRoom.players = append(gameRoom.players, *player)
}

func (gr *GameRoom) listen() {
	for {
		select {
		case c := <-gr.
		}
	}
}