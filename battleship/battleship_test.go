package battleship

import "testing"

func Test_Create_Game(t *testing.T) {
	game := NewGame(0,1)
	if game.players[0].Id != 0 {
		t.Errorf("game did not create player1 with id 0")
	}
	if game.players[1].Id != 1 {
		t.Errorf("game did not create player1 with id 0")
	}
}

func Test_Add_Ship(t *testing.T) {
	game := NewGame(0,1)
	game.PlaceShip(0,0,0)
	p1 := &game.players[0]
	p2 := &game.players[1]
	if p2.Shoot(p1, 0,0) == false {
		t.Errorf("player1 does not have ship at 0,0")
	}
}

func Test_Game_FIGHT(t *testing.T) {
	game := NewGame(0,1)
	game.PlaceShip(1,0,0)
	myTurn, isHit := game.FIGHT(0,0,0,)
	if myTurn == false {
		t.Errorf("player turn error")
	}
	if isHit == false {
		t.Errorf("player not hit after FIGHT")
	}

	game = NewGame(0,1)
	game.PlaceShip(1,0,0)
	myTurn, isHit = game.FIGHT(0,1,1,)
	if myTurn == false {
		t.Errorf("player turn error")
	}
	if isHit == true {
		t.Errorf("player hit after FIGHT")
	}
	if game.turn != 1 {
		t.Errorf("not changing turns")
	}

	game = NewGame(0,1)
	game.PlaceShip(1,0,0)
	myTurn, isHit = game.FIGHT(1,1,1,)
	if myTurn == true {
		t.Errorf("player turn error")
	}
	if isHit == true {
		t.Errorf("player hit after FIGHT")
	}
}