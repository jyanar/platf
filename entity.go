package main

import "github.com/hajimehoshi/ebiten/v2"

// An entity has a physical location, can update, and has an associated
// draw method. It also has a notify action, which can be used to relay
// messages elsewhere.
type Entity interface {
	PositionalObject
	notify(msg string)
	Update() error
	Draw(screen *ebiten.Image)
}
