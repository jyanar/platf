package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Various utilities functions

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func sign(x float64) float64 {
	if math.Signbit(x) {
		return -1.0
	} else {
		return 1.0
	}
}

func moveToward(start, stop, step float64) float64 {
	switch {
	case start < stop:
		if start+step >= stop {
			return stop
		}
		return start + step

	case start > stop:
		if start-step <= stop {
			return stop
		}
		return start - step

	default:
		return stop
	}
}

func drawGrid(screen *ebiten.Image) {
	for i := 0; i < WINDOWSIZE; i += 16 {
		for j := 0; j < WINDOWSIZE; j += 16 {
			ebitenutil.DrawRect(screen, float64(i), float64(j), 1, 1, color.White)
		}
	}
}

type Vector struct {
	x, y float64
}

func add(a, b Vector) Vector {
	return Vector{a.x + b.x, a.y + b.y}
}

func subtract(a, b Vector) Vector {
	return Vector{a.x - b.x, a.y - b.y}
}

func dot(a, b Vector) float64 {
	return a.x*b.x + a.y*b.y
}
