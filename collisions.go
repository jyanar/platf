package main

import "fmt"

// Collisions is a collection of objects with positional data
type Collisions struct {
	items []*Obj
}

func NewCollisions() *Collisions {
	c := &Collisions{}
	c.items = []*Obj{}
	return c
}

func (c *Collisions) init() {
	c.items = []*Obj{}
}

func (c *Collisions) add(item *Obj) {
	c.items = append(c.items, item)
}

func (w Collisions) areOverlapping(a, b *Obj) bool {
	aX, aY, aW, aH := a.getPosAndSize()
	bX, bY, bW, bH := b.getPosAndSize()
	return aX < bX+bW && // aX is less than b+width
		   aX+aW > bX && // a+width is greater than bX
		   aY < bY+bH && // a is less than b+height
		   aH+aY > bY    // but, a+height is greater than b
}

func (c Collisions) computeOverlap(a, b *Obj) (width float64, height float64) {
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

func (c Collisions) checkIsColliding(item *Obj) (collidingobj *Obj) {
	for _, other := range c.items {
		other := other
		if other != item && other.isSolid && c.areOverlapping(item, other) {
			collidingobj = other
		}
	}
	return
}

func (c *Collisions) move(item *Obj, newX float64, newY float64) {
	prevX, prevY := item.x, item.y

	// move in the x and check collisions
	item.x = newX
	if colObj := c.checkIsColliding(item); colObj != nil {
		width, _ := c.computeOverlap(item, colObj)
		if newX > prevX {
			newX = newX - width // item is moving left
		} else {
			newX = newX + width // item is moving right
		}
		item.x = newX
	}
	// move in the y and check collisions
	item.y = newY
	if colObj := c.checkIsColliding(item); colObj != nil {
		fmt.Printf("COLLIDING WITH: \n")
		fmt.Println(colObj)
		_, height := c.computeOverlap(item, colObj)
		if newY > prevY {
			newY = newY - height // item is moving left
		} else {
			newY = newY + height // item is moving right
		}
		// item.setPosition(newX, newY)
		item.y = newY
	}
}

// func (c Collisions) getItem(item string) Entity {
// 	for i := range c.items {
// 		if typeof(i) == "*main.Lever" {
// 			return i
// 		}
// 	}
// 	return nil
// }

func (c Collisions) printAllItems() {
	for i := range c.items {
		fmt.Println(c.items[i])
	}
}