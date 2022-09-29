package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	World
	Player
	Level
}

func (g *Game) Update() error {
	g.Player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.Draw(screen)
	g.Level.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WINDOWSIZE, WINDOWSIZE
}

func (g *Game) init() {
	g.World.init()
	g.Player.init(10, 10, &g.World)
	g.Level.init()
	g.World.add(&g.Player)
	for i := 0; i < len(g.Level.tiles); i++ {
		g.World.add(&g.Level.tiles[i])
	}
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
