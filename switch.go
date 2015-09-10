package main

import "github.com/nsf/termbox-go"

type Direction uint

const (
	UP    Direction = iota
	RIGHT           = iota
	DOWN            = iota
	LEFT            = iota
)

var directions = []Direction{UP, RIGHT, DOWN, LEFT}

type Rotation uint

const (
	CLOCKWISE         Rotation = iota
	COUNTER_CLOCKWISE          = iota
)

// a swinging door or set of doors around the same pivot
// center
type Switch struct {
	grid     Grid
	pivot    Position
	position Position // position within the level
	rotation Direction
}

func NewSwitch() *Switch {
	return &Switch{Grid{
		{false, false, true},
		{false, false, true},
		{false, false, false, true, true},
	},
		Position{2, 2}, Position{5, 5}, UP}
}

func (s *Switch) Draw() {

	plp := s.position.Add(s.pivot)
	termbox.SetCell(plp.X, plp.Y, 'O', foregroundColor, backgroundColor)

	for y := 0; y < len(s.grid); y++ {
		for x := 0; x < len(s.grid[y]); x++ {
			if s.grid[y][x] {
				var r rune
				if x > s.pivot.X {
					r = '-'
				} else if x < s.pivot.X {
					r = '-'
				} else if y > s.pivot.Y {
					r = '|'
				} else if y < s.pivot.Y {
					r = '|'
				} else {
					panic("bad door position")
				}
				gridp := s.position.Add(Position{x, y})
				termbox.SetCell(gridp.X, gridp.Y, r, foregroundColor, backgroundColor)
			}
		}
	}
}

/*
func (s *Switch) Collides(p Position) bool {
}

func (s *Switch) CollideSwivel(oldp Position, newp Position) Rotation {
}
*/
