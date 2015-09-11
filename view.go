package main

import (
	"github.com/nsf/termbox-go"
)

const backgroundColor = termbox.ColorBlack
const foregroundColor = termbox.ColorWhite
const wallColor = termbox.ColorBlue

func render() {
	termbox.Clear(backgroundColor, backgroundColor)
	level.drawOutside()
	level.drawWalls()
	level.drawTreasure()
	level.drawExit()
	level.drawSwitches()
	level.drawMonsters()
	hero.Draw()
	termbox.Flush()
}
