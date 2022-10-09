package main

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	Obj
	*Collisions
	vy    float64
	speed float64
	alive bool
}

func NewPlayer(obj Obj, c *Collisions, vy float64, speed float64, alive bool) *Player {
	return &Player{
		Obj:        obj,
		Collisions: c,
		vy:         0,
		speed:      220,
		alive:      true,
	}
}

func (p Player) NewGroundedObj() *Obj {
	return &Obj{p.x, p.y + p.h, p.w, 1}
}

func (p Player) NewHittingCeilingObj() *Obj {
	return &Obj{p.x, p.y - 1, p.w, 1}
}

// func (p *Player) init(obj Obj, w *Collisions, vy float64, speed float64, alive bool) *Player {
// 	return &Player{
// 		Obj:   obj,
// 		collisions: w,
// 		vy:    0,
// 		speed: 220,
// 		alive: true,
// 	}
// }

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

func (p Player) isGrounded() bool {
	return p.Collisions.checkIsColliding(p.NewGroundedObj()) != nil
}

func (p Player) isHittingCeiling() bool {
	return p.Collisions.checkIsColliding(p.NewHittingCeilingObj()) != nil
}

func (p *Player) checkDead() {
	if obj := p.Collisions.checkIsColliding(p.NewGroundedObj()); obj != nil && typeof(obj) == "*main.Spikes" {
		p.alive = false
	}
}

func (p *Player) Update() error {
	p.alive = true
	dt := 0.01
	dx := 0.0
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && p.isGrounded() {
		p.vy = JUMP_SPEED
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		dx = dx - dt*p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		dx = dx + dt*p.speed
	}
	p.Collisions.move(p, p.x+dx, p.y+p.vy)
	p.vy = math.Min(p.vy+GRAVITY*dt, 12)
	if p.isGrounded() {
		p.vy = 0
	}
	if p.isHittingCeiling() {
		p.vy = GRAVITY * dt
	}
	p.checkDead()
	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, p.x, p.y, p.w, p.h, image.White)
}
