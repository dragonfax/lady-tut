package main

import "github.com/nsf/termbox-go"

type Grid [][]bool

type Level struct {
	width  uint
	height uint
	walls  Walls
	// switches []Switch
	monsters []Monster
}

type Walls Grid

var level = &Level{
	10,
	10,
	Walls{
		{false, false, false, false, false, true},
		{false, false, false, false, true},
	},
	// []Switch{},
	[]Monster{Monster{Position{5, 5}}},
}

func (l *Level) drawOutside() {
	w, h := termbox.Size()
	for x := 0; x <= w; x++ {
		for y := 0; y <= h; y++ {
			if uint(x) < l.width && uint(y) < l.height {
				termbox.SetCell(x, y, ' ', backgroundColor, backgroundColor)
			} else {
				termbox.SetCell(x, y, ' ', wallColor, wallColor)
			}
		}
	}
}

func (l *Level) drawWalls() {
	for y := 0; y < len(l.walls); y++ {
		for x := 0; x < len(l.walls[y]); x++ {
			if l.walls[y][x] {
				termbox.SetCell(x, y, ' ', wallColor, wallColor)
			}
		}
	}
}

func (l *Level) isWallAt(p Position) bool {
	if p.X < 0 || uint(p.X) >= l.width || p.Y < 0 || uint(p.Y) >= l.height {
		return true
	}
	if len(l.walls) > p.Y && len(l.walls[p.Y]) > p.X && l.walls[p.Y][p.X] {
		return true
	}
	return false
}

// level.drawMonsters()
