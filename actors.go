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
	h.position.X -= 1
	if h.position.X < 0 {
		h.position.X = 0
	}
}

func (h *Hero) moveRight() {
	h.position.X += 1
}

func (h *Hero) moveUp() {
	h.position.Y -= 1
	if h.position.Y < 0 {
		h.position.Y = 0
	}
}

func (h *Hero) moveDown() {
	h.position.Y += 1
}
