package main

type Tile struct {
	Pos
	w, h int
}

func (t Tile) getPosition() (float64, float64) {
	return t.X, t.Y
}

func (t *Tile) setPosition(X, Y float64) {
	t.X = X
	t.Y = Y
}

func (t Tile) getPosAndSize() (float64, float64, int, int) {
	return t.Pos.X, t.Pos.Y, t.w, t.h
}
