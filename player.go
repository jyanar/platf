package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	Pos
	world     *World
	idx, w, h int
	speed     float64
}

func (p *Player) setIdx(idx int) {
	p.idx = idx
}

func (p Player) getIdx() int {
	return p.idx
}

func (p *Player) setPosition(X, Y float64) {
	p.X = X
	p.Y = Y
}

func (p Player) getPosition() (float64, float64) {
	return p.X, p.Y
}

func (p Player) getPosAndSize() (float64, float64, int, int) {
	return p.Pos.X, p.Pos.Y, p.w, p.h
}

func (p *Player) init(X float64, Y float64, w *World) {
	p.X = X
	p.Y = Y
	p.world = w
	p.w = TILESIZE
	p.h = TILESIZE
	p.speed = 3
}

func (p *Player) Update() error {
	dx, dy := 0.0, 0.0
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		dy += p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		dy -= p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		dx -= p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		dx += p.speed
	}
	p.world.move(p, p.X+dx, p.Y+dy)
	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, p.X, p.Y, float64(p.w), float64(p.h), image.White)
}
