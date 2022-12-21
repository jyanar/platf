package main

import (
	"fmt"
)

// Collisions is a collection of Objects with positional, velocity,
// solid, and width/height data
type Collisions struct {
	objects []*Object
}

func NewCollisions() *Collisions {
	c := &Collisions{}
	c.objects = []*Object{}
	return c
}

func (c *Collisions) init() {
	c.objects = []*Object{}
}

func (c *Collisions) add(item *Object) {
	c.objects = append(c.objects, item)
}

func (c *Collisions) areOverlapping(a, b *Object) bool {
	aX, aY, aW, aH := a.getPosAndSize()
	bX, bY, bW, bH := b.getPosAndSize()
	return aX < bX+bW && // aX is less than b+width
		aX+aW > bX && // a+width is greater than bX
		aY < bY+bH && // a is less than b+height
		aH+aY > bY // but, a+height is greater than b
}

func (c *Collisions) computeOverlap(a, b *Object) (width float64, height float64) {
	ax, ay, aw, ah := a.getPosAndSize()
	bx, by, bw, bh := b.getPosAndSize()
	if ax+aw > bx+bw {
		width = bx + bw - ax
	} else {
		width = ax + bw - bx
	}
	if ay+ah > by+bh {
		height = by + bh - ay
	} else {
		height = ay + ah - by
	}
	return width, height
}

func (c *Collisions) checkIsColliding(item *Object) (collidingobj *Object) {
	for _, other := range c.objects {
		other := other
		if other != item && other.isSolid && c.areOverlapping(item, other) {
			collidingobj = other
		}
	}
	return
}

func (c *Collisions) integrateVelocity(object *Object) {

	prevX, prevY := object.position.x, object.position.y

	// Integrate velocity in x direction
	object.position.x = object.position.x + object.velocity.x*DT
	if colObj := c.checkIsColliding(object); colObj != nil {
		// If this is a platform, ignore
		if !colObj.isPlatform {
			width, _ := c.computeOverlap(object, colObj)
			if object.position.x > prevX {
				object.position.x -= width // object is moving left
			} else {
				object.position.x += width // object is moving right
			}
		}
	}

	// Integrate velocity in y direction
	collisionflag := false
	object.position.y = object.position.y + object.velocity.y
	if colObj := c.checkIsColliding(object); colObj != nil {
		collisionflag = true
		if colObj.isPlatform && object.velocity.y < 0 {
			collisionflag = false
		}
		if collisionflag {
			_, height := c.computeOverlap(object, colObj)
			if object.position.y > prevY {
				object.position.y -= height
			} else {
				object.position.y += height
			}
		}
	}

	// // If this is a platform, ignore if we're jumping
	// switch {
	// case object.velocity.y < 0:
	// 	if !colObj.isPlatform {
	// 		_, height := c.computeOverlap(object, colObj)
	// 		if object.position.y > prevY {
	// 			object.position.y -= height // object is moving left
	// 		} else {
	// 			object.position.y += height // object is moving right
	// 		}
	// 	}
	// case object.velocity.y > 0: // falling
	// 	_, height := c.computeOverlap(object, colObj)
	// 	if object.position.y > prevY {
	// 		object.position.y -= height // object is moving left
	// 	} else {
	// 		object.position.y += height // object is moving right
	// 	}
	// }
	// }
}

func (c *Collisions) move(object *Object, newX float64, newY float64) {
	prevX, prevY := object.position.x, object.position.y

	// move in the x and check collisions
	object.position.x = newX
	if colObj := c.checkIsColliding(object); colObj != nil {
		width, _ := c.computeOverlap(object, colObj)
		if newX > prevX {
			newX = newX - width // object is moving left
		} else {
			newX = newX + width // object is moving right
		}
		object.position.x = newX
	}
	// move in the y and check collisions
	object.position.y = newY
	if colObj := c.checkIsColliding(object); colObj != nil {
		_, height := c.computeOverlap(object, colObj)
		if newY > prevY {
			newY = newY - height // item is moving left
		} else {
			newY = newY + height // item is moving right
		}
		object.position.y = newY
	}
}

// func (c Collisions) getItem(item string) Entity {
// 	for i := range c.objects {
// 		if typeof(i) == "*main.Lever" {
// 			return i
// 		}
// 	}
// 	return nil
// }

func (c Collisions) printAllobjects() {
	for i := range c.objects {
		fmt.Println(c.objects[i])
	}
}
