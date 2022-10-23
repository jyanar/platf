package main

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jyanar/platf/graphics"
)

type Enemy struct {
	Obj
	*Collisions
	velocity Vector
	speed    float64
	alive    bool
	dir      int
	anim     string
	anims    map[string]*graphics.Animation
}

func NewEnemy(obj Obj, col *Collisions) *Enemy {
	enemy := Enemy{
		Obj:        obj,
		Collisions: col,
		velocity:   Vector{0, 0},
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

func (e Enemy) NewGroundedObj() *Obj {
	return &Obj{e.x, e.y + e.h, e.w, 1, true}
}

func (e Enemy) NewHittingCeilingObj() *Obj {
	return &Obj{e.x, e.y - 1, e.w, 1, true}
}

func (e Enemy) NewSideObj() *Obj {
	if e.dir == -1 {
		return &Obj{e.x - 1, e.y, 1, e.h / 2, true}
	} else {
		return &Obj{e.x + 16, e.y, 1, e.h / 2, true}
	}
}

func (e *Enemy) setPosition(x, y float64) {
	e.x = x
	e.y = y
}

func (e Enemy) getPosition() (float64, float64) {
	return e.x, e.y
}

func (e Enemy) getPosAndSize() (float64, float64, float64, float64) {
	return e.x, e.y, e.w, e.h
}

func (e Enemy) isGrounded() bool {
	return e.Collisions.checkIsColliding(e.NewGroundedObj()) != nil
}

func (e Enemy) isHittingCeiling() bool {
	return e.Collisions.checkIsColliding(e.NewHittingCeilingObj()) != nil
}

func (e Enemy) willHitWall() bool {
	return e.Collisions.checkIsColliding(e.NewSideObj()) != nil
}

func (e Enemy) willFall() bool {
	grobj := e.NewGroundedObj()
	if e.dir == -1 {
		grobj.x -= 16
		return e.checkIsColliding(grobj) == nil
	} else {
		grobj.x += 16
		return e.checkIsColliding(grobj) == nil
	}
}

func (e Enemy) Solid() bool {
	return e.Obj.Solid()
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

	e.Collisions.move(&e.Obj, e.x+e.velocity.x*dt, e.y+e.velocity.y)

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
		op.GeoM.Translate(e.x+16-4, e.y)
	} else {
		op.GeoM.Translate(e.x-4, e.y)
	}
	screen.DrawImage(graphics.Quads[e.anims[e.anim].GetFrame()], op)
	// Debugging collisiosn
	// e.NewSideObj().Draw(screen)
	// e.Obj.Draw(screen)
}
