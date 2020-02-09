package ship

type Element interface {
	isHit(x int, y int)
}

type Ship struct {
	X int
	Y int
	hit bool
}

func NewShip(x int, y int) Ship {
	var s Ship
	s.X = x
	s.Y = y
	return s
}

func (s *Ship) isHit(x int, y int) bool{
	if s.X == x && s.Y == y {
		s.hit = true
	}
	return s.hit
}