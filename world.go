package main

// World is a colletion of objects with positional data
type World struct {
	items []PosObj
}

func (w *World) init() {
	w.items = []PosObj{}
}

func (w *World) add(item PosObj) {
	w.items = append(w.items, item)
}

func (w World) areOverlapping(a, b PosObj) bool {
	aX, aY, aW, aH := a.getPosAndSize()
	bX, bY, bW, bH := b.getPosAndSize()
	return aX < bX + bW && // aX is less than b+width
		   aX + aW > bX && // a+width is greater than bX
		   aY < bY + bH && // a is less than b+height
		   aH + aY > bY    // but, a+height is greater than b
}

func (w World) computeOverlap(a, b PosObj) (width float64, height float64) {
	ax, ay, aw, ah := a.getPosAndSize()
	bx, by, bw, bh := b.getPosAndSize()
	if ax+aw > bx+bw {
		width = bx+bw - ax
	} else {
		width = ax+bw - bx
	}
	if ay+ah > by+bh {
		height = by+bh - ay
	} else {
		height = ay+ah - by
	}
	return width, height
}

func (w World) checkIsColliding(item PosObj) (collidingobj PosObj) {
	for _, other := range w.items {
		if other != item && w.areOverlapping(item, other) {
			collidingobj = other
		}
	}
	return
}

func (w *World) move(item PosObj, newX float64, newY float64) {
	prevX, prevY := item.getPosition()

	// move in the x and check collisions
	item.setPosition(newX, prevY)
	if colObj := w.checkIsColliding(item); colObj != nil {
		width, _ := w.computeOverlap(item, colObj)
		if newX > prevX {
			newX = newX - width // item is moving left
		} else {
			newX = newX + width // item is moving right
		}
		item.setPosition(newX, prevY)
	}
	// move in the y and check collisions
	item.setPosition(newX, newY)
	if colObj := w.checkIsColliding(item); colObj != nil {
		_, height := w.computeOverlap(item, colObj)
		if newY > prevY {
			newY = newY - height // item is moving left
		} else {
			newY = newY + height // item is moving right
		}
		item.setPosition(newX, newY)
	}
}
