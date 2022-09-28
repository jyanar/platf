package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Pos struct {
	X, Y float64
}

type PosObj interface {
	setIdx(idx int)
	getIdx() int
	getPosition() (float64, float64)
	setPosition(X, Y float64)
	getPosAndSize() (float64, float64, int, int)
}

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
	for _, t := range g.Level.tiles {
		g.World.add(&t)
	}
	for i, t := range g.World.items {
		fmt.Printf("%v %v\n", i, t.getIdx())
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
