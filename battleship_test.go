package battleship

import "testing"

func Test_Create_Game(t *testing.T) {
	game := NewGame()
	if game.Player1.Id != 1 {
		t.Errorf("game did not create player1 with id 0")
	}
	if game.Player2.Id != 2 {
		t.Errorf("game did not create player1 with id 0")
	}
}
