// main file for battleship
// James Rine & Matt Perry

package battleship

import . "github.com/Jlrine2/CS372_Battleship/battleship/player"

type Game struct {
	Player1 Player
	Player2 Player
}

func NewGame() Game {
	var game Game
	game.Player1 = NewPlayer(1)
	game.Player2 = NewPlayer(2)
	return game
}
