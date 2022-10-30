package main

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jyanar/platf/graphics"
)

type Enemy struct {
	Object
	*Collisions
	speed float64
	alive bool
	dir   int
	anim  string
	anims map[string]*graphics.Animation
}

func NewEnemy(obj Object, col *Collisions) *Enemy {
	enemy := Enemy{
		Object:     obj,
		Collisions: col,
		speed:      220,
		alive:      true,
		dir:        1,
		anim:       "idle",
		anims:      make(map[string]*graphics.Animation),
	}
	enemy.anims["idle"] = graphics.NewAnimation(20, 4, 0.5)
	enemy.anims["run"] = graphics.NewAnimation(20, 4, 0.5)
	return &enemy
}

func (e Enemy) NewGroundedObject() *Object {
	return &Object{
		position: Vector{e.position.x, e.position.y + e.h},
		velocity: Vector{0, 0},
		w:        e.w,
		h:        1,
		isSolid:  true,
	}
}

func (e Enemy) NewHittingCeilingObject() *Object {
	return &Object{
		position: Vector{e.position.x, e.position.y - 1},
		w:        e.w,
		h:        1,
		isSolid:  true,
	}
}

func (e Enemy) NewSideObject() *Object {
	if e.dir == -1 {
		return &Object{
			position: Vector{e.position.x - 1, e.position.y},
			w:        1,
			h:        e.h / 2,
			isSolid:  true,
		}
	} else {
		return &Object{
			position: Vector{e.position.x + 16, e.position.y},
			w:        1,
			h:        e.h / 2,
			isSolid:  true,
		}
	}
}

func (e Enemy) isGrounded() bool {
	return e.Collisions.checkIsColliding(e.NewGroundedObject()) != nil
}

func (e Enemy) isHittingCeiling() bool {
	return e.Collisions.checkIsColliding(e.NewHittingCeilingObject()) != nil
}

func (e Enemy) willHitWall() bool {
	return e.Collisions.checkIsColliding(e.NewSideObject()) != nil
}

func (e Enemy) willFall() bool {
	grObject := e.NewGroundedObject()
	if e.dir == -1 {
		grObject.position.x -= 16
		return e.checkIsColliding(grObject) == nil
	} else {
		grObject.position.x += 16
		return e.checkIsColliding(grObject) == nil
	}
}

func (e Enemy) Solid() bool {
	return e.Object.isSolid
}

func (e *Enemy) Update(state *GameState) error {
	// log.Printf("UPDATING ENEMY")
	dt := 0.1

	// Randomly decide to move, if sitting still
	switch e.anim {
	case "idle":
		directions := []int{-1, 1}
		e.dir = directions[rand.Intn(len(directions))]
		e.anim = "run"
	case "run":
		switch e.dir {
		case -1:
			e.velocity.x -= ACCELERATION / 2

		case 1:
			e.velocity.x += ACCELERATION / 2
		}
	}

	// Limit to MAX_SPEED / MAX_FALL_SPEED
	if math.Abs(e.velocity.x) > MAX_SPEED {
		e.velocity.x = sign(e.velocity.x) * MAX_SPEED
	}
	e.velocity.y = math.Min(e.velocity.y+GRAVITY*dt, MAX_FALL_SPEED)

	if e.willHitWall() || e.willFall() {
		e.dir = -1 * e.dir
	}

	e.Collisions.integrateVelocity(&e.Object)
	// e.Collisions.move(&e.Object, e.x+e.velocity.x*dt, e.y+e.velocity.y)

	if e.isGrounded() {
		e.velocity.y = 0
	}
	if e.isHittingCeiling() {
		e.velocity.y = GRAVITY * dt
	}

	// Apply friction
	if e.velocity.x > 0 {
		e.velocity.x = moveToward(e.velocity.x, 0, FRICTION)
	} else if e.velocity.x < 0 {
		e.velocity.x = moveToward(e.velocity.x, 0, FRICTION)
	}

	// Update animation
	e.anims[e.anim].Update()
	return nil
}

func (e Enemy) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if e.dir == -1 {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(e.position.x+16-4, e.position.y)
	} else {
		op.GeoM.Translate(e.position.x-4, e.position.y)
	}
	screen.DrawImage(graphics.Quads[e.anims[e.anim].GetFrame()], op)
	// Debugging collisiosn
	// e.NewSideObject().Draw(screen)
	// e.Object.Draw(screen)
}
