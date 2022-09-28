package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Pos struct {
	X, Y float64
}

type Game struct {
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
	g.Player.init()
	g.Level.init()
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
