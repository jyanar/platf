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
	lastdir  int
	anim     string
	anims    map[string]*graphics.Animation
}

func NewPlayer(obj Obj, col *Collisions) *Player {
	player := Player{
		Obj:        obj,
		Collisions: col,
		velocity:   Vector{0, 0},
		speed:      220,
		alive:      true,
		lastdir:    1,
		anim:       "idle",
		anims:      make(map[string]*graphics.Animation),
	}
	player.anims["idle"] = graphics.NewAnimation(8, 3, 0.5)
	player.anims["run"] = graphics.NewAnimation(12, 7, 0.6)
	player.anims["jump"] = graphics.NewAnimation(12, 3, 0.2)
	return &player
}

func (p Player) NewGroundedObj() *Obj {
	return &Obj{p.x, p.y + p.h, p.w, 1, true}
}

func (p Player) NewHittingCeilingObj() *Obj {
	return &Obj{p.x, p.y - 1, p.w, 1, true}
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
	// Set default animation to idle
	p.anim = "idle"

	// Take in input
	dt := 0.1
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && p.isGrounded() {
		p.anim = "jump"
		p.velocity.y = JUMP_SPEED
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.anim = "run"
		p.lastdir = -1
		p.velocity.x -= ACCELERATION
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.anim = "run"
		p.lastdir = 1
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

	// Update animation
	p.anims[p.anim].Update()
	return nil
}

func (p Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if p.lastdir == -1 {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(p.x+16-4, p.y)
	} else {
		op.GeoM.Translate(p.x-4, p.y)
	}
	screen.DrawImage(graphics.Quads[p.anims[p.anim].GetFrame()], op)
	// p.anim.Draw(screen)
	// screen.DrawImage(graphics.Player, op)
}
