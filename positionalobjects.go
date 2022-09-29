package main

type Obj struct {
	x, y, w, h float64
}

type PosObj interface {
	getPosition() (float64, float64)
	setPosition(x, y float64)
	getPosAndSize() (float64, float64, float64, float64)
}

func (o *Obj) setPosition(x, y float64) {
	o.x = x
	o.y = y
}

func (o Obj) getPosition() (float64, float64) {
	return o.x, o.y
}

func (o Obj) getPosAndSize() (float64, float64, float64, float64) {
	return o.x, o.y, o.w, o.h
}
