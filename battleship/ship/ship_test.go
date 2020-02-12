package ship

import "testing"

func Test_Ship_Init(t *testing.T) {
	var s Ship
	if s.X != 0 {
		t.Error("new ship doesn't initalize X coord")
	}
	if s.Y != 0 {
		t.Error("new ship doesn't initalize Y coord")
	}
	if s.hit != false {
		t.Error("new ship doesn't initalize hit to false")
	}
}

func Test_Ship_Construction(t *testing.T) {
	s := NewShip(1,1)
	if s.X != 1 {
		t.Errorf("new ship doesn't initalize X coord to 1, got %d", s.X)
	}
	if s.Y != 1 {
		t.Errorf("new ship doesn't initalize X coord to 1, got %d", s.Y)
	}
	if s.hit != false {
		t.Error("new ship doesn't initalize hit to false")
	}
}

func Test_Shooting(t *testing.T) {
	s1 := NewShip(0,0)
	s1.isHit(0,0)
	if s1.hit != true {
		t.Errorf("shot at ship didn't update hit value %t", s1.hit)
	}
	s2 := NewShip(8,8)
	s2.isHit(8,8)
	if s2.hit != true {
		t.Errorf("shot at ship didn't update hit value %t", s2.hit)
	}
}