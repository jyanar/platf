package main

import (
	"math"
)

type Level struct {
	width int
	data  []int
	// tiles []Tile
}

func (l *Level) init(world *World) {
	m := map[int]string{
		1: "Tile",
		2: "ToggleFloor",
		5: "Spikes",
		6: "Lever",
		9: "Player",
	}

	l.width = 16
	l.data = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 2, 2, 2, 2, 2, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		9, 0, 0, 6, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 1, 1, 5, 5, 1, 1,
	}
	for i := 0; i < len(l.data); i++ {
		x, y := float64((i % l.width) * TILESIZE),  math.Floor(float64(i)/float64(l.width)) * TILESIZE
		obj  := Obj{x: x, y: y, w: float64(TILESIZE), h: float64(TILESIZE)}
		switch m[l.data[i]] {
		case "Tile":
			world.add(&Tile{obj})

		case "ToggleFloor":
			world.add(&ToggleFloor{obj})

		case "Spikes":
			world.add(&Spikes{obj})

		case "Lever":
			world.add(&Lever{obj})

		case "Player":
			world.add(&Player{obj, world, 0, 220})
		}

		// if l.data[i] > 0 {
		// 	t := Tile{Obj{
		// 		x: float64(float64(i%l.width) * TILESIZE),
		// 		y: float64(math.Floor(float64(i/l.width)) * TILESIZE),
		// 		w: float64(TILESIZE),
		// 		h: float64(TILESIZE),
		// 	}}
		// 	l.tiles = append(l.tiles, t)
		// }
	}
}

// func (l *Level) Draw(screen *ebiten.Image) {
// 	for _, t := range l.tiles {
// 		ebitenutil.DrawRect(screen, t.x, t.y, t.w, t.h, image.White)
// 	}
// }
