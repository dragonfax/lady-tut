package main

type Position struct {
	X int
	Y int
}

func (p Position) Add(q Position) Position {
	return Position{p.X + q.X, p.Y + q.Y}
}
