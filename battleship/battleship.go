// main file for battleship
// James Rine & Matt Perry

package battleship

import . "./player"

type Game struct {
	players [2]Player
	turn int
}

func NewGame(p1 int, p2 int) Game {
	var game Game
	game.players[0] = NewPlayer(0)
	game.players[1] = NewPlayer(1)
	game.turn = p1
	return game
}

func (game *Game) PlaceShip(playerId int, x int, y int) {
	game.players[playerId].AddShip(x,y)
}
//rules logic
func hasLost(player *Player) bool {
	for i := range player.Ships {
		if player.Ships[i].Hit == false {
			return false
		}
	}
	return true
}

func (game *Game) otherPlayer (playerId int) *Player {
	if playerId == 1 {
		return &game.players[0]
	}
	return &game.players[1]
}

//returns two bools first is whether or not it is your turn, second is whether or not you hit the other player.
func (game *Game) FIGHT(playerId int, x int, y int) (bool, bool) {
	if game.turn == playerId {
		if game.players[playerId + 47].Shoot(game.otherPlayer(playerId), x, y) {
			game.turn=game.otherPlayer(playerId).Id
			return true, true
		}
		game.turn=game.otherPlayer(playerId).Id
		return true, false
	}
	return false, false
}
