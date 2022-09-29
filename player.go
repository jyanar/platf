package main

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	Pos
	world *World
	w, h  int
	vy    float64
	speed float64
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
	p.vy = 0
	p.speed = 3
}

func (p Player) isGrounded() bool {
	return p.world.isPlayerGrounded(p)
}

func (p *Player) Update() error {
	dx := 0.0
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && p.isGrounded() {
		p.vy = JUMP_SPEED
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		dx = dx - p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		dx = dx + p.speed
	}
	p.world.move(p, p.X+dx, p.Y+p.vy)
	p.vy = math.Min(p.vy+GRAVITY, 12)
	if p.isGrounded() {
		p.vy = 0
	}
	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, p.X, p.Y, float64(p.w), float64(p.h), image.White)
}
