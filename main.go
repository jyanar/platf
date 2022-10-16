package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jyanar/platf/graphics"
)

func main() {
	ebiten.SetWindowSize(WINDOWSIZE*2, WINDOWSIZE*2)
	ebiten.SetWindowTitle("platf")
	sm := &SceneManager{}
	sm.push(&PlayScene{})
	err := graphics.Load()
	if err != nil {
		fmt.Println("ERROR!")
	}
	if err := ebiten.RunGame(sm); err != nil {
		log.Fatal(err)
	}
}
