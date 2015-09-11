package main

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

type Hero struct {
	position  Position
	health    uint
	treasures uint
}

var hero = Hero{Position{0, 0}, 1, 0}

func (h Hero) Draw() {
	termbox.SetCell(h.position.X, h.position.Y, 'H', foregroundColor, backgroundColor)
}

func (h Hero) collides(np Position, op Position) bool {

	if level.exit == np {
		termbox.Close()
		fmt.Printf("Win! Treasure = %d\n", h.treasures*10)
		os.Exit(0)
	}

	if level.isWallAt(np) {
		return true
	}

	if level.isSwitchAt(np, op, true) {
		return true
	}

	return false
}

func (h *Hero) takeTreasure() {
	for i, t := range level.treasure {
		if t != nil && *t == h.position {
			h.treasures++
			level.treasure[i] = nil
		}
	}
}

func (h *Hero) moveLeft() {
	np := h.position

	np.X -= 1
	if np.X < 0 {
		np.X = 0
	}

	if h.collides(np, h.position) {
		return
	}

	h.takeTreasure()

	h.position = np
}

func (h *Hero) moveRight() {
	np := h.position

	np.X += 1

	if h.collides(np, h.position) {
		return
	}

	h.takeTreasure()

	h.position = np
}

func (h *Hero) moveUp() {
	np := h.position

	np.Y -= 1
	if np.Y < 0 {
		np.Y = 0
	}

	if h.collides(np, h.position) {
		return
	}

	h.takeTreasure()

	h.position = np
}

func (h *Hero) moveDown() {
	np := h.position

	np.Y += 1

	if h.collides(np, h.position) {
		return
	}

	h.takeTreasure()

	h.position = np
}
