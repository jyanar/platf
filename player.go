package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jyanar/platf/graphics"
)

type Player struct {
	Obj
	*Collisions
	velocity Vector
	speed    float64
	alive    bool
}

func NewPlayer(obj Obj, c *Collisions) *Player {
	return &Player{
		Obj:        obj,
		Collisions: c,
		velocity:   Vector{0, 0},
		speed:      220,
		alive:      true,
	}
}

func (p Player) NewGroundedObj() *Obj {
	return &Obj{p.x, p.y + p.h, p.w, 1, true}
}

func (p Player) NewHittingCeilingObj() *Obj {
	return &Obj{p.x, p.y - 1, p.w, 1, true}
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

func (p Player) Solid() bool {
	return p.Obj.Solid()
}

func (p *Player) Update(state *GameState) error {
	// Take in input
	dt := 0.1
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && p.isGrounded() {
		p.velocity.y = JUMP_SPEED
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.velocity.x -= ACCELERATION
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.velocity.x += ACCELERATION
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		state.SceneManager.getCurrent().trigger("player:action")
	}

	// Limit to MAX_SPEED / MAX_FALL_SPEED
	if math.Abs(p.velocity.x) > MAX_SPEED {
		p.velocity.x = sign(p.velocity.x) * MAX_SPEED
	}
	p.velocity.y = math.Min(p.velocity.y+GRAVITY*dt, MAX_FALL_SPEED)

	p.Collisions.move(&p.Obj, p.x+p.velocity.x*dt, p.y+p.velocity.y)

	if p.isGrounded() {
		p.velocity.y = 0
	}
	if p.isHittingCeiling() {
		p.velocity.y = GRAVITY * dt
	}

	// Apply friction
	if p.velocity.x > 0 {
		p.velocity.x = moveToward(p.velocity.x, 0, FRICTION)
	} else if p.velocity.x < 0 {
		p.velocity.x = moveToward(p.velocity.x, 0, FRICTION)
	}
	return nil
}

func moveToward(start, stop, step float64) float64 {
	switch {
	case start < stop:
		if start+step >= stop {
			return stop
		}
		return start + step

	case start > stop:
		if start-step <= stop {
			return stop
		}
		return start - step

	default:
		return stop
	}
}

func (p Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.x-4, p.y)
	screen.DrawImage(graphics.Player, op)
}
