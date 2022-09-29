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

func (w World) checkCollisions(a, b PosObj) bool {
	aX, aY, aW, aH := a.getPosAndSize()
	bX, bY, bW, bH := b.getPosAndSize()
	return aX < bX+float64(bW) && aX+float64(aW) > bX && aY < bY+float64(bH) && float64(aH)+aY > bY
}

func (w World) checkIsColliding(item PosObj) bool {
	for _, other := range w.items {
		if other != item && w.checkCollisions(item, other) {
			return true
		}
	}
	return false
}

func (w World) isPlayerGrounded(p Player) bool {
	// construct an item right below the player.
	iObj := IObj{Pos{p.X, p.Y + float64(p.h)}, p.w, 8}
	return w.checkIsColliding(&iObj)
}

func (w *World) move(item PosObj, newX float64, newY float64) {
	prevX, prevY := item.getPosition()
	item.setPosition(newX, newY)
	if w.checkIsColliding(item) {
		item.setPosition(prevX, prevY)
	}
}

// func (w *World) move(idx int, newX float64, newY float64) bool {
// 	prevX, prevY := w.items[idx].getPosition()
// 	w.items[idx].setPosition(newX, newY)
// 	for itr, item := range w.items {
// 		if itr != idx && w.checkCollisions(w.items[idx], item) {
// 			w.items[idx].setPosition(prevX, prevY)
// 			return false
// 		}
// 	}
// 	return true
// }
