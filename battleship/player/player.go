package player

import (
	"github.com/Jlrine2/CS372_Battleship/battleship/ship"
)

type Player struct {
	id    int
	ships []ship.Ship
}

func NewPlayer(id int) Player {
	var p Player
	p.id = id
	return p
}

func (p *Player) AddShip(x int, y int) {
	p.ships = append(p.ships, ship.NewShip(x, y))
}

func (p *Player) Shoot(opponent *Player, x int, y int) bool {
	for i := range opponent.ships {
		if opponent.ships[i].IsHit(x, y) == true {
			return true
		}
	}
	return false
}
