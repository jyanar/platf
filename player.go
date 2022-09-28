package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	Pos
	w, h  int
	speed float64
}

func (p *Player) init() {
	p.Pos.X = 10.0
	p.Pos.Y = 10.0
	p.w = TILESIZE
	p.h = TILESIZE
	p.speed = 3
}

func (p *Player) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Pos.Y = p.Pos.Y + p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Pos.Y = p.Pos.Y - p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.Pos.X = p.Pos.X - p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.Pos.X = p.Pos.X + p.speed
	}
	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, p.X, p.Y, float64(p.w), float64(p.h), image.White)
}
