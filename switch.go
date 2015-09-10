package main

import "github.com/nsf/termbox-go"

type Direction uint8

// if the default door is facing up
// what direction will it face after its been rotated?
const (
	UP            Direction = iota
	RIGHT                   = iota
	DOWN                    = iota
	LEFT                    = iota
	MAX_DIRECTION           = iota
)

type Rotation int8

const (
	CLOCKWISE         Rotation = 1
	COUNTER_CLOCKWISE          = -1
)

// a swinging door or set of doors around the same pivot
// center
type Switch struct {
	grid      Grid
	pivot     Position
	position  Position // position within the level
	direction Direction
}

func NewSwitch() *Switch {
	return &Switch{
		Grid{
			{false, false, true},
			{false, false, true},
			{false, false, false, true, true},
		},
		Position{2, 2},
		Position{5, 5},
		UP,
	}
}

// given a rotated position, return the unrotated/original position
func (s *Switch) UnwindPosition(p Position) Position {
	o := p.Subtract(s.pivot)

	var no Position

	switch s.direction {
	case UP:
		return p
	case RIGHT:
		no = Position{o.Y, -o.X}
	case DOWN:
		no = Position{-o.X, -o.Y}
	case LEFT:
		no = Position{-o.Y, o.X}
	default:
		panic("unknown case")
	}

	return no.Add(s.pivot)
}

// given a position in the switch grid, return the rotated position.
func (s *Switch) WindPosition(p Position) Position {
	// 1. orient to pivot point
	// 2. flip axis and sign
	// 3. reorient to pivot again

	o := p.Subtract(s.pivot)

	var no Position

	switch s.direction {
	case UP:
		return p
	case RIGHT:
		no = Position{-o.Y, o.X}
	case DOWN:
		no = Position{-o.X, -o.Y}
	case LEFT:
		no = Position{o.Y, -o.X}
	default:
		panic("unknown case")
	}

	return no.Add(s.pivot)
}

func (s *Switch) Draw() {

	plp := s.position.Add(s.pivot)
	termbox.SetCell(plp.X, plp.Y, 'O', foregroundColor, backgroundColor)

	for y := 0; y < len(s.grid); y++ {
		for x := 0; x < len(s.grid[y]); x++ {
			if s.grid[y][x] {
				rp := s.WindPosition(Position{x, y})
				var r rune
				if rp.X > s.pivot.X {
					r = '-'
				} else if rp.X < s.pivot.X {
					r = '-'
				} else if rp.Y > s.pivot.Y {
					r = '|'
				} else if rp.Y < s.pivot.Y {
					r = '|'
				} else {
					panic("bad door position")
				}
				gridp := s.position.Add(rp)
				termbox.SetCell(gridp.X, gridp.Y, r, foregroundColor, backgroundColor)
			}
		}
	}
}

func (s *Switch) isCollided(p0 Position) (bool, bool) {
	p1 := p0.Subtract(s.position)
	p := s.UnwindPosition(p1)
	if p.X < 0 || p.Y < 0 {
		return false, false
	}
	if p == s.pivot {
		return true, false
	}
	if len(s.grid) > p.Y && len(s.grid[p.Y]) > p.X && s.grid[p.Y][p.X] {
		return true, true
	}
	return false, false
}

func (s *Switch) Swivel(op Position) {
	// TODO
	s.Rotate(CLOCKWISE)
}

func (s *Switch) Rotate(c Rotation) {
	nd := uint8(s.direction) + uint8(c)
	if nd >= MAX_DIRECTION {
		nd = 0
	}
	if nd < 0 {
		nd = MAX_DIRECTION - 1
	}
	s.direction = Direction(nd)
}
