package main

import (
	"bufio"
	"os"

	"github.com/nsf/termbox-go"
)

type Grid [][]bool

type Level struct {
	width    uint
	height   uint
	walls    Grid
	switches []*Switch
	monsters []*Monster
	exit     Position
	entrance Position
	treasure []*Position
}

var level *Level

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

func (l *Level) isSwitchAt(p Position, op Position, flip bool) bool {
	for _, s := range l.switches {
		if a := s.isCollided(p); a {
			if can, r := s.canRotate(op); flip && can {
				s.rotate(r)
			}
			return true
		}
	}
	return false
}

func (l *Level) drawSwitches() {
	for _, s := range l.switches {
		s.Draw()
	}
}

func (l *Level) drawMonsters() {
	for _, m := range l.monsters {
		m.Draw()
	}
}

func load(filename string) *Level {
	l := &Level{walls: make([][]bool, 0, 50)}

	file, err := os.Open(filename)
	if err != nil {
		panic("failed to open level file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, 50)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic("scanner error: " + err.Error())
	}

	if len(lines) == 0 {
		panic("no lines read from file")
	}

	// scan for individual sells
	for y, line := range lines {
		for x, b := range line {
			if b == 'W' {
				if y >= len(l.walls) {
					l.walls = append(l.walls, make([]bool, 0, 50))
				}
				for x >= len(l.walls[y]) {
					l.walls[y] = append(l.walls[y], false)
				}
				l.walls[y][x] = true
			}

			if b == 'H' {
				l.entrance = Position{x, y}
			}

			if b == 'E' {
				l.exit = Position{x, y}
			}

			if b == 'T' {
				l.treasure = append(l.treasure, &Position{x, y})
			}

			if b == 'M' {
				l.monsters = append(l.monsters, NewMonster(x, y))
			}

			if b == 'O' {
				l.switches = append(l.switches, &Switch{position: Position{x, y}})
			}
		}

		if len(line) > int(l.width) {
			l.width = uint(len(line))
		}
	}
	l.height = uint(len(lines))

	// scan for switch panels
	for _, s := range l.switches {

		if lines[s.position.Y-1][s.position.X] == '|' {
			s.panels[NORTH] = true

			// measure panel width
			for i := 1; lines[s.position.Y-i][s.position.X] == '|'; i++ {
				s.width = i
			}
		}

		if lines[s.position.Y+1][s.position.X] == '|' {
			s.panels[SOUTH] = true

			// measure panel width
			for i := 1; lines[s.position.Y+i][s.position.X] == '|'; i++ {
				s.width = i
			}
		}

		if lines[s.position.Y][s.position.X-1] == '-' {
			s.panels[WEST] = true

			// measure panel width
			for i := 1; lines[s.position.Y][s.position.X-i] == '-'; i++ {
				s.width = int(i / 2)
			}
		}

		if lines[s.position.Y][s.position.X+1] == '-' {
			s.panels[EAST] = true

			// measure panel width
			for i := 1; lines[s.position.Y][s.position.X+i] == '-'; i++ {
				s.width = int(i / 2)
			}
		}

	}

	return l
}
