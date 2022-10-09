package main

// World is a colletion of objects with positional data
type CollisionSystem struct {
	items []PositionalObject
}

func NewCollisionSystem() *CollisionSystem {
	cs := &CollisionSystem{}
	cs.items = []PositionalObject{}
	return cs
}

func (cs *CollisionSystem) init() {
	cs.items = []PositionalObject{}
}

func (cs *CollisionSystem) add(item PositionalObject) {
	cs.items = append(cs.items, item)
}

func (cs CollisionSystem) areOverlapping(a, b PositionalObject) bool {
	aX, aY, aW, aH := a.getPosAndSize()
	bX, bY, bW, bH := b.getPosAndSize()
	return aX < bX+bW && // aX is less than b+width
		aX+aW > bX && // a+width is greater than bX
		aY < bY+bH && // a is less than b+height
		aH+aY > bY // but, a+height is greater than b
}

func (cs CollisionSystem) computeOverlap(a, b PositionalObject) (width float64, height float64) {
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

// func (cs CollisionSystem) checkIsColliding(item PositionalObject) (objects []PositionalObject) {
// 	for _, other := range cs.items {
// 		if other != item && cs.areOverlapping(item, other) {
// 			objects = append(objects, other)
// 		}
// 	}
// }

func (cs CollisionSystem) checkIsColliding(item PositionalObject) (collidingobj PositionalObject) {
	for _, other := range cs.items {
		if other != item && cs.areOverlapping(item, other) {
			collidingobj = other
		}
	}
	return
}

func (cs *CollisionSystem) move(item PositionalObject, newX float64, newY float64) {
	prevX, prevY := item.getPosition()

	// move in the x and check collisions
	item.setPosition(newX, prevY)
	if colObj := cs.checkIsColliding(item); colObj != nil {
		width, _ := cs.computeOverlap(item, colObj)
		if newX > prevX {
			newX = newX - width // item is moving left
		} else {
			newX = newX + width // item is moving right
		}
		item.setPosition(newX, prevY)
	}
	// move in the y and check collisions
	item.setPosition(newX, newY)
	if colObj := cs.checkIsColliding(item); colObj != nil {
		_, height := cs.computeOverlap(item, colObj)
		if newY > prevY {
			newY = newY - height // item is moving left
		} else {
			newY = newY + height // item is moving right
		}
		item.setPosition(newX, newY)
	}
}

// func (cs *CollisionSystem) Update() error {
// 	for _, item := range w.items {
// 		item.Update()
// 	}
// 	return nil
// }

// func (cs CollisionSystem) Draw(screen *ebiten.Image) {
// 	for _, item := range w.items {
// 		item.Draw(screen)
// 	}
// }

// func (cs CollisionSystem) getPlayer() *Player {
// 	for _, item := range w.items {
// 		if typeof(item) == "*main.Player" {
// 			return &item
// 		}
// 	}
// 	return &Player{}
// }
