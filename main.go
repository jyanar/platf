package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	World
	Player
	Level
}

func (g *Game) Update() error {
	// g.Player.Update()
	g.World.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// g.Player.Draw(screen)
	// g.Level.Draw(screen)
	g.World.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WINDOWSIZE, WINDOWSIZE
}

func (g *Game) init() {
	g.World.init()
	g.Level.init(&g.World)
	// g.Player.init(10, 10, &g.World)
	// g.Level.init()
	// g.World.add(&g.Player)
	// for i := 0; i < len(g.Level.tiles); i++ {
	// 	g.World.add(&g.Level.tiles[i])
	// }
}

func NewGame() *Game {
	g := &Game{}
	g.init()
	return g
}

func main() {
	ebiten.SetWindowSize(WINDOWSIZE, WINDOWSIZE)
	ebiten.SetWindowTitle("platf")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
