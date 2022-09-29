package main

type Pos struct {
	X, Y float64
}

type PosObj interface {
	getPosition() (float64, float64)
	setPosition(X, Y float64)
	getPosAndSize() (float64, float64, int, int)
}

type IObj struct {
	Pos
	w, h int
}

func (i IObj) init(X float64, Y float64, w int, h int) IObj {
	return IObj{Pos{X, Y}, w, h}
}

func (i *IObj) setPosition(X, Y float64) {
	i.X = X
	i.Y = Y
}

func (i IObj) getPosition() (float64, float64) {
	return i.X, i.Y
}

func (i IObj) getPosAndSize() (float64, float64, int, int) {
	return i.X, i.Y, i.w, i.h
}
