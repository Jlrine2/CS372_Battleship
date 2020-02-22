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
	if s.Hit != false {
		t.Error("new ship doesn't initalize Hit to false")
	}
}

func Test_Ship_Construction(t *testing.T) {
	s := NewShip(1, 1)
	if s.X != 1 {
		t.Errorf("new ship doesn't initalize X coord to 1, got %d", s.X)
	}
	if s.Y != 1 {
		t.Errorf("new ship doesn't initalize X coord to 1, got %d", s.Y)
	}
	if s.Hit != false {
		t.Error("new ship doesn't initalize Hit to false")
	}

	s1 := NewShip(19, 0)
	if s1.X != 19 {
		t.Errorf("new ship doesn't initalize X coord to 1, got %d", s1.X)
	}
	if s1.Y != 0 {
		t.Errorf("new ship doesn't initalize X coord to 1, got %d", s1.Y)
	}
	if s1.Hit != false {
		t.Error("new ship doesn't initalize Hit to false")
	}

	s2 := NewShip(100, 99)
	if s2.X != 100 {
		t.Errorf("new ship doesn't initalize X coord to 1, got %d", s2.X)
	}
	if s2.Y != 99 {
		t.Errorf("new ship doesn't initalize X coord to 1, got %d", s2.Y)
	}
	if s2.Hit != false {
		t.Error("new ship doesn't initalize Hit to false")
	}
}

func Test_Shooting(t *testing.T) {
	s1 := NewShip(0, 0)
	s1.IsHit(0, 0)
	if s1.Hit != true {
		t.Errorf("shot at ship didn't update Hit value %t", s1.Hit)
	}
	s2 := NewShip(8, 8)
	s2.IsHit(8, 8)
	if !s2.Hit {
		t.Errorf("shot at ship didn't update Hit value %t", s2.Hit)
	}
	s3 := NewShip(9, 0)
	s3.IsHit(9, 0)
	if !s3.Hit {
		t.Errorf("shot at ship didn't update Hit value %t", s3.Hit)
	}
	s4 := NewShip(800, 12)
	s4.IsHit(800, 12)
	if !s4.Hit {
		t.Errorf("shot at ship didn't update Hit value %t", s4.Hit)
	}
	s5 := NewShip(-8, 8)
	s5.IsHit(-8, 8)
	if !s5.Hit {
		t.Errorf("shot at ship didn't update Hit value %t", s5.Hit)
	}
	s6 := NewShip(1111, 1111)
	s6.IsHit(1111, 1111)
	if !s6.Hit {
		t.Errorf("shot at ship didn't update Hit value %t", s6.Hit)
	}

}
