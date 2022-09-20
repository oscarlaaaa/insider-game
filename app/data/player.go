package data

import (
	"log"

	"github.com/gorilla/websocket"
)

// Game Roles
type Role int64

const (
	Master     = "master"
	Insider    = "insider"
	Common     = "common"
	Unselected = "unselected"
)

type Player struct {
	Id   int             `json:"playerId"`
	Name string          `json:"playerName"`
	Conn *websocket.Conn `json:"-"`
	Send chan []byte     `json:"-"`
	Room *GameRoom       `json:"-"`
}

func (p *Player) ReadPump() {
	for {
		messageType, data, err := p.Conn.ReadMessage()
		if err != nil {

		}
		switch messageType {
		case websocket.TextMessage:
			log.Println(string(data))
		default:
			log.Println("no idea")
		}
	}
}

func (p *Player) subscribe() {}
