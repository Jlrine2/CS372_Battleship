package ship

type Ship struct {
	X   int
	Y   int
	Hit bool
}

func NewShip(x int, y int) Ship {
	var s Ship
	s.X = x
	s.Y = y
	return s
}

func (s *Ship) IsHit(x int, y int) bool {
	if s.X == x && s.Y == y {
		s.Hit = true
		return true
	}
	return false
}
