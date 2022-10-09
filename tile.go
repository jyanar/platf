package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tile struct {
	// TODO I think it potentially makes more sense to put the specific parent state here, rather than the SceneManager?
	// However, you should not construct a pointer to an interface (SceneManager is a struct)
	sm *SceneManager
	Obj
}

func NewTile(sm *SceneManager, obj Obj) *Tile {
	return &Tile{sm, obj}
}

func (t Tile) notify(msg string) {
	t.sm.getCurrent().processMsg(msg)
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

func (t Tile) Update() error { return nil }

func (t Tile) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, t.x, t.y, t.w, t.h, image.White)
}
