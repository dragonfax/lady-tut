package main

type Direction uint

const (
	UP    Direction = iota
	RIGHT           = iota
	DOWN            = iota
	LEFT            = iota
)

var directions = []Direction{UP, RIGHT, DOWN, LEFT}

// a swinging door or set of doors around the same pivot
// center
type Switch struct {
	grid     Grid
	pivot    Position
	position Position // position within the level
	rotation Direction
}

/*
func (s *Switch) Rotate() {
	i := uint(s.rotation) + 1
	if i >= len(directions) {
		i = 0
	}
	s.rotation = directions[i]
}
*/
