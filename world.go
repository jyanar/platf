package main

// World is a colletion of objects with positional data
type World struct {
	items []PosObj
}

func (w *World) init() {
	w.items = []PosObj{}
}

func (w *World) add(item PosObj) {
	item.setIdx(len(w.items))
	w.items = append(w.items, item)
}

func (w *World) checkCollisions(a, b PosObj) bool {
	aX, aY, aW, aH := a.getPosAndSize()
	bX, bY, bW, bH := b.getPosAndSize()
	return aX < bX+float64(bW) && aX+float64(aW) > bX &&
		aY < bY+float64(bH) && float64(aH)+aY > bY
}

func (w *World) move(item PosObj, newX float64, newY float64) {
	prevX, prevY := item.getPosition()
	item.setPosition(newX, newY)
	for _, i := range w.items {
		if i != item && w.checkCollisions(item, i) {
			item.setPosition(prevX, prevY)
			return
		}
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
