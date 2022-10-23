package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jyanar/platf/graphics"
)

func main() {
	ebiten.SetWindowSize(WINDOWSIZE*WINDOWSCALE, WINDOWSIZE*WINDOWSCALE)
	ebiten.SetWindowTitle("platf")
	sm := &SceneManager{}
	sm.push(&PlayScene{})
	err := graphics.Load()
	if err != nil {
		fmt.Println(err)
	}
	if err := ebiten.RunGame(sm); err != nil {
		log.Fatal(err)
	}
}
