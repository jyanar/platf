package main

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Level struct {
	width int
	data  []int
	tiles []Tile
}

func (l *Level) init() {
	l.width = 16
	l.data = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0,
		1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1,
	}
	for i := 0; i < len(l.data); i++ {
		if l.data[i] > 0 {
			t := Tile{Obj{
				x: float64(float64(i%l.width) * TILESIZE),
				y: float64(math.Floor(float64(i/l.width)) * TILESIZE),
				w: float64(TILESIZE),
				h: float64(TILESIZE),
			}}
			l.tiles = append(l.tiles, t)
		}
	}
}

func (l *Level) Draw(screen *ebiten.Image) {
	for _, t := range l.tiles {
		ebitenutil.DrawRect(screen, t.x, t.y, t.w, t.h, image.White)
	}
}
