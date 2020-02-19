package player

import (
	"../ship"
)

type Player struct {
	Id    int
	Ships []ship.Ship
}

func NewPlayer(id int) Player {
	var p Player
	p.Id = id
	return p
}

func (p *Player) AddShip(x int, y int) {
	p.Ships = append(p.Ships, ship.NewShip(x, y))
}

func (p *Player) Shoot(opponent *Player, x int, y int) bool {
	for i := range opponent.Ships {
		if opponent.Ships[i].IsHit(x, y) == true {
			return true
		}
	}
	return false
}
