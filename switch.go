package main

import "github.com/nsf/termbox-go"

type Panel int

const (
	NORTH Panel = iota
	EAST
	SOUTH
	WEST
	MAX_PANEL int = 4
)

type Rotation int8

const (
	CLOCKWISE         Rotation = 1
	COUNTER_CLOCKWISE Rotation = -1
)

// a swinging door or set of doors around the same pivot
// center
type Switch struct {
	width    int             // door length, doubled for east/west doors
	position Position        // coord of pivot within the level
	rotation int             // 0-3
	panels   [MAX_PANEL]bool // starting with the north panel
}

func NewSwitch() *Switch {
	return &Switch{
		width:    2,
		position: Position{5, 5},
		rotation: 0,
		panels:   [4]bool{true, true, false, false},
	}
}

func (s *Switch) hasPanel(p Panel) bool {

	rp := (int(p) - s.rotation) % int(MAX_PANEL)

	// reverse modulus
	if rp < 0 {
		rp = 4 + rp
	}

	return s.panels[rp]
}

func (s *Switch) Draw() {

	termbox.SetCell(s.position.X, s.position.Y, 'O', foregroundColor, backgroundColor)

	if s.hasPanel(NORTH) {
		for i := int(1); i <= s.width; i++ {
			termbox.SetCell(s.position.X, s.position.Y-i, '|', foregroundColor, backgroundColor)
		}
	}

	if s.hasPanel(EAST) {
		for i := 1; i <= s.width*2; i++ {
			termbox.SetCell(s.position.X+i, s.position.Y, '-', foregroundColor, backgroundColor)
		}
	}

	if s.hasPanel(SOUTH) {
		for i := 1; i <= s.width; i++ {
			termbox.SetCell(s.position.X, s.position.Y+i, '|', foregroundColor, backgroundColor)
		}
	}

	if s.hasPanel(WEST) {
		for i := 1; i <= s.width*2; i++ {
			termbox.SetCell(s.position.X-i, s.position.Y, '-', foregroundColor, backgroundColor)
		}
	}

}

func (s *Switch) isCollided(p Position) bool {

	if p.X == s.position.X && p.Y == s.position.Y {
		// collides with center
		return true
	}

	if s.hasPanel(NORTH) && p.X == s.position.X && p.Y < s.position.Y && p.Y >= s.position.Y-s.width {
		return true
	}

	if s.hasPanel(SOUTH) && p.X == s.position.X && p.Y > s.position.Y && p.Y <= s.position.Y+s.width {
		return true
	}

	if s.hasPanel(WEST) && p.Y == s.position.Y && p.X < s.position.X && p.X >= s.position.X-s.width*2 {
		return true
	}

	if s.hasPanel(EAST) && p.Y == s.position.Y && p.X > s.position.X && p.X <= s.position.X+s.width*2 {
		return true
	}

	return false
}

func (s *Switch) canRotate(p Position) (bool, Rotation) {

	if s.hasPanel(NORTH) && p.Y < s.position.Y && p.Y >= s.position.Y-s.width {
		if p.X == s.position.X+1 {
			return true, COUNTER_CLOCKWISE
		} else if p.X == s.position.X-1 {
			return true, CLOCKWISE
		}
	}

	if s.hasPanel(SOUTH) && p.Y > s.position.Y && p.Y <= s.position.Y+s.width {
		if p.X == s.position.X+1 {
			return true, CLOCKWISE
		} else if p.X == s.position.X-1 {
			return true, COUNTER_CLOCKWISE
		}
	}

	if s.hasPanel(WEST) && p.X < s.position.X && p.X >= s.position.X-s.width*2 {
		if p.Y == s.position.Y+1 {
			return true, CLOCKWISE
		} else if p.Y == s.position.Y-1 {
			return true, COUNTER_CLOCKWISE
		}
	}

	if s.hasPanel(EAST) && p.X > s.position.X && p.X <= s.position.X+s.width*2 {
		if p.Y == s.position.Y+1 {
			return true, COUNTER_CLOCKWISE
		} else if p.Y == s.position.Y-1 {
			return true, CLOCKWISE
		}
	}

	return false, CLOCKWISE
}

func (s *Switch) rotate(r Rotation) {
	s.rotation = s.rotation + int(r)
	if s.rotation == 4 {
		s.rotation = 0
	}
	if s.rotation == -1 {
		s.rotation = 3
	}
}
