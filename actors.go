package main

import "github.com/nsf/termbox-go"

type Position struct {
	X int
	Y int
}

type Hero struct {
	position Position
	health   uint
}

type Monster struct {
	position Position
}

var hero = Hero{Position{0, 0}, 1}

func (h Hero) Draw() {
	termbox.SetCell(h.position.X, h.position.Y, 'H', foregroundColor, backgroundColor)
}

func (h *Hero) moveLeft() {
	np := h.position

	np.X -= 1
	if np.X < 0 {
		np.X = 0
	}

	if level.isWallAt(np) {
		return
	}

	h.position = np
}

func (h *Hero) moveRight() {
	np := h.position

	np.X += 1

	if level.isWallAt(np) {
		return
	}

	h.position = np
}

func (h *Hero) moveUp() {
	np := h.position

	np.Y -= 1
	if np.Y < 0 {
		np.Y = 0
	}

	if level.isWallAt(np) {
		return
	}

	h.position = np
}

func (h *Hero) moveDown() {
	np := h.position

	np.Y += 1

	if level.isWallAt(np) {
		return
	}

	h.position = np
}
