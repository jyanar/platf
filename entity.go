package main

import "github.com/hajimehoshi/ebiten/v2"

type Entity interface {
	PositionalObject
	Solid() bool
	Update(state *GameState) error
	Draw(screen *ebiten.Image)
}
