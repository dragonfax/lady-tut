package main

import (
	"time"

	"github.com/nsf/termbox-go"
)

func main() {

	termbox.Init()
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Key == termbox.KeyArrowLeft:
					hero.moveLeft()
				case ev.Key == termbox.KeyArrowRight:
					hero.moveRight()
				case ev.Key == termbox.KeyArrowUp:
					hero.moveUp()
				case ev.Key == termbox.KeyArrowDown:
					hero.moveDown()
				case ev.Key == termbox.KeyEsc:
					return
				}
			}
		// case <-monsterMoveTimer.C:
		// level.MoveMonsters()
		default:
			render()
			time.Sleep(10 * time.Millisecond)
		}
	}
}
