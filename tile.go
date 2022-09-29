package main

type Tile struct {
	Obj
}

func (t Tile) getPosition() (float64, float64) {
	return t.x, t.y
}

func (t *Tile) setPosition(x, y float64) {
	t.x = x
	t.y = y
}

func (t Tile) getPosAndSize() (float64, float64, float64, float64) {
	return t.x, t.y, t.w, t.h
}
