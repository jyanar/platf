package main

// Collisions is a colletion of objects with positional data
type Collisions struct {
	items []PosObj
}

func NewCollisions() *Collisions {
	c := &Collisions{}
	c.items = []PosObj{}
	return c
}

func (c *Collisions) init() {
	c.items = []PosObj{}
}

func (c *Collisions) add(item PosObj) {
	c.items = append(c.items, item)
}

func (w Collisions) areOverlapping(a, b PosObj) bool {
	aX, aY, aW, aH := a.getPosAndSize()
	bX, bY, bW, bH := b.getPosAndSize()
	return aX < bX+bW && // aX is less than b+width
		aX+aW > bX && // a+width is greater than bX
		aY < bY+bH && // a is less than b+height
		aH+aY > bY // but, a+height is greater than b
}

func (c Collisions) computeOverlap(a, b PosObj) (width float64, height float64) {
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

func (c Collisions) checkIsColliding(item PosObj) (collidingobj PosObj) {
	for _, other := range c.items {
		if other != item && c.areOverlapping(item, other) {
			collidingobj = other
		}
	}
	return
}

func (c *Collisions) move(item PosObj, newX float64, newY float64) {
	prevX, prevY := item.getPosition()

	// move in the x and check collisions
	item.setPosition(newX, prevY)
	if colObj := c.checkIsColliding(item); colObj != nil {
		width, _ := c.computeOverlap(item, colObj)
		if newX > prevX {
			newX = newX - width // item is moving left
		} else {
			newX = newX + width // item is moving right
		}
		item.setPosition(newX, prevY)
	}
	// move in the y and check collisions
	item.setPosition(newX, newY)
	if colObj := c.checkIsColliding(item); colObj != nil {
		_, height := c.computeOverlap(item, colObj)
		if newY > prevY {
			newY = newY - height // item is moving left
		} else {
			newY = newY + height // item is moving right
		}
		item.setPosition(newX, newY)
	}
}

// func (w *Collisions) Update() error {
// 	for _, item := range w.items {
// 		item.Update()
// 	}
// 	return nil
// }

// func (w Collisions) Draw(screen *ebiten.Image) {
// 	for _, item := range w.items {
// 		item.Draw(screen)
// 	}
// }

// func (w Collisions) getPlayer() *Player {
// 	for _, item := range w.items {
// 		if typeof(item) == "*main.Player" {
// 			return &item
// 		}
// 	}
// 	return &Player{}
// }
