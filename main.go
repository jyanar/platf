package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(WINDOWSIZE*2, WINDOWSIZE*2)
	ebiten.SetWindowTitle("platf")
	sm := &StateManager{}
	sm.push(&PlayState{})
	// sm.setCurrent(&PlayState{})
	if err := ebiten.RunGame(sm); err != nil {
		log.Fatal(err)
	}
}
