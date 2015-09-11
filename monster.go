package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

type Monster struct {
	position Position
}

func NewMonster(x, y int) *Monster {
	m := &Monster{Position{x, y}}
	go m.ThinkLoop()
	return m
}

func (m Monster) Draw() {
	termbox.SetCell(m.position.X, m.position.Y, 'M', foregroundColor, backgroundColor)
}

func (m Monster) collides(np Position) bool {
	if level.isWallAt(np) {
		return true
	}

	if level.isSwitchAt(np, Position{}, false) {
		return true
	}

	return false
}

func (m *Monster) moveLeft() {
	np := m.position

	np.X -= 1
	if np.X < 0 {
		np.X = 0
	}

	if hero.position == np {
		termbox.Close()
		os.Exit(1)
	}

	if m.collides(np) {
		return
	}

	m.position = np
}

func (m *Monster) moveUp() {
	np := m.position

	np.Y -= 1
	if np.Y < 0 {
		np.Y = 0
	}

	if hero.position == np {
		termbox.Close()
		os.Exit(1)
	}

	if m.collides(np) {
		return
	}

	m.position = np
}

func (m *Monster) moveRight() {
	np := m.position

	np.X += 1

	if hero.position == np {
		termbox.Close()
		os.Exit(1)
	}

	if m.collides(np) {
		return
	}

	m.position = np
}

func (m *Monster) moveDown() {
	np := m.position

	np.Y += 1

	if hero.position == np {
		termbox.Close()
		os.Exit(1)
	}

	if m.collides(np) {
		return
	}

	m.position = np
}

func (m *Monster) Think() {
	// if I can see the Hero, move toward the hero
	// otherwise move randomly

	if hero.position.X == m.position.X {
		if hero.position.Y > m.position.Y {
			m.moveDown()
		} else {
			m.moveUp()
		}
	} else if hero.position.Y == m.position.Y {
		if hero.position.X > m.position.X {
			m.moveRight()
		} else {
			m.moveLeft()
		}
	} else {
		[]func(){m.moveUp, m.moveDown, m.moveLeft, m.moveRight}[rand.Uint32()%4]()
	}
}

// a goroutine to think and move at a specific pace.
func (m *Monster) ThinkLoop() {
	ticker := time.NewTicker(time.Millisecond * 500)
	for {
		select {
		case <-ticker.C:
			m.Think()
		}
	}
}
