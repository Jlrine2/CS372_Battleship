package player

import (
	"testing"
)

func Test_Player_Init(t *testing.T) {
	var p Player
	if p.Id != 0 {
		t.Error("new player doesn't initalize Id")
	}
}

func Test_Player_Construction(t *testing.T) {
	p1 := NewPlayer(1)
	if p1.Id != 1 {
		t.Error("new player doesn't initalize Id")
	}
	p2 := NewPlayer(123456789)
	if p2.Id != 123456789 {
		t.Error("new player doesn't initalize Id")
	}
}

func Test_Player_Ships(t *testing.T) {
	p := NewPlayer(1)
	if len(p.Ships) != 0 {
		t.Error("Player Ships not empty on initialization")

	}
}

func Test_Player_Add_Ships(t *testing.T) {
	p := NewPlayer(1)
	p.AddShip(1, 1)
	if len(p.Ships) != 1 {
		t.Error("Player Ships not added")
	}
	s := &p.Ships[0]
	if s.X != 1 && s.Y != 1 {
		t.Errorf("New Ship not as constructed: Expect x=1,y=1: Got x=%d, y=%d", s.X, s.Y)
	}
	p.AddShip(1, 2)
	if len(p.Ships) != 2 {
		t.Errorf("Player Ships not added")
	}
	s = &p.Ships[0]
	if s.X != 1 && s.Y != 1 {
		t.Errorf("New Ship not as constructed: Expect x=1,y=1: Got x=%d, y=%d", s.X, s.Y)
	}
	s = &p.Ships[1]
	if s.X != 1 && s.Y != 2 {
		t.Errorf("New Ship not as constructed: Expect x=1,y=1: Got x=%d, y=%d", s.X, s.Y)
	}

}

func Test_Player_Shoot(t *testing.T) {
	p1 := NewPlayer(1)
	p2 := NewPlayer(2)
	p2.AddShip(1, 1)
	p1.Shoot(&p2, 1, 1)
	s1 := &p2.Ships[0]
	if s1.Hit != true {
		t.Logf("ship at: x=%d, y=%d", s1.X, s1.Y)
		t.Errorf("ship not set to hit after recieving shot: %t", s1.Hit)
	}
	p2.AddShip(4, 4)
	p1.Shoot(&p2, 4, 4)
	s2 := &p2.Ships[1]
	if s2.Hit != true {
		t.Logf("ship at: x=%d, y=%d", s2.X, s2.Y)
		t.Errorf("ship not set to hit after recieving shot: %t", s2.Hit)
	}
}
