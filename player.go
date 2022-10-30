package main

import (
	// "log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jyanar/platf/graphics"
)

type Player struct {
	Object
	*Collisions
	alive   bool
	lastdir int
	anim    string
	anims   map[string]*graphics.Animation
}

func NewPlayer(Object Object, col *Collisions) *Player {
	player := Player{
		Object:     Object,
		Collisions: col,
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

func (p Player) NewGroundedObject() *Object {
	return &Object{
		position: Vector{p.position.x, p.position.y + p.h},
		w:        p.w,
		h:        1,
		isSolid:  true,
	}
}

func (p Player) NewHittingCeilingObject() *Object {
	return &Object{
		position: Vector{p.position.x, p.position.y - 1},
		w:        p.w,
		h:        1,
		isSolid:  true,
	}
}

func (p Player) NewSideObject() *Object {
	if p.lastdir == -1 {
		return &Object{
			position: Vector{p.position.x - 1, p.position.y},
			w:        1,
			h:        p.h / 2,
			isSolid:  true,
		}
	} else {
		return &Object{
			position: Vector{p.position.x + 16, p.position.y},
			w:        1,
			h:        p.h / 2,
			isSolid:  true,
		}
	}
}

func (p Player) isGrounded() bool {
	return p.Collisions.checkIsColliding(p.NewGroundedObject()) != nil
}

func (p Player) isHittingCeiling() bool {
	return p.Collisions.checkIsColliding(p.NewHittingCeilingObject()) != nil
}

func (p Player) Solid() bool {
	return p.Object.isSolid
}

func (p *Player) Update(state *GameState) error {
	// Set default animation to idle
	p.anim = "idle"

	// Take in input
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) && p.isGrounded() {
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
	p.velocity.y = math.Min(p.velocity.y+GRAVITY*DT, MAX_FALL_SPEED)

	p.Collisions.integrateVelocity(&p.Object)
	// p.Collisions.move(&p.Object, p.x+p.velocity.x*dt, p.y+p.velocity.y)

	if p.isGrounded() {
		p.velocity.y = 0
	}
	if p.isHittingCeiling() {
		p.velocity.y = GRAVITY * DT
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
		op.GeoM.Translate(p.position.x+16-4, p.position.y)
	} else {
		op.GeoM.Translate(p.position.x-4, p.position.y)
	}
	/* log.Printf("Player pos: %v", p.Object) */
	screen.DrawImage(graphics.Quads[p.anims[p.anim].GetFrame()], op)
	/* ebitenutil.DrawRect(screen, p.x, p.y, p.w, p.h, color.White) */
}
