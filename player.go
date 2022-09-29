package main

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	Obj
	world *World
	vy    float64
	speed float64
}

func (p *Player) setPosition(x, y float64) {
	p.x = x
	p.y = y
}

func (p Player) getPosition() (float64, float64) {
	return p.x, p.y
}

func (p Player) getPosAndSize() (float64, float64, float64, float64) {
	return p.x, p.y, p.w, p.h
}

func (p *Player) init(x float64, y float64, w *World) {
	p.Obj = Obj{x, y, TILESIZE, TILESIZE}
	p.world = w
	p.vy = 0
	p.speed = 220
}

func (p Player) isGrounded() bool {
	return p.world.checkIsColliding(&Obj{p.x, p.y + p.h, p.w, 1}) != nil
}

func (p *Player) Update() error {
	dt := 0.01
	dx := 0.0
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && p.isGrounded() {
		p.vy = JUMP_SPEED
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		dx = dx - dt * p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		dx = dx + dt * p.speed
	}
	p.world.move(p, p.x+dx, p.y+p.vy)
	p.vy = math.Min(p.vy + GRAVITY * dt, 12)
	if p.isGrounded() {
		p.vy = 0
	}
	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, p.x, p.y, p.w, p.h, image.White)
}
