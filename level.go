package main

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tile struct {
	x, y, w, h int
}

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
			t := Tile{
				x: (i % l.width) * TILESIZE,
				y: int(math.Floor(float64(i/l.width)) * TILESIZE),
				w: TILESIZE,
				h: TILESIZE,
			}
			l.tiles = append(l.tiles, t)
		}
	}
}

// func init() {
// 	fmt.Println(l.tiles)
// }

func (l *Level) Draw(screen *ebiten.Image) {
	for _, t := range l.tiles {
		// ebitenutil.DrawLine(screen, float64(t.x), float64(t.y), float64(t.x + t.w), float64(t.y + t.h), image.White)
		ebitenutil.DrawRect(screen, float64(t.x), float64(t.y), float64(t.w), float64(t.h), image.White)
	}
}
