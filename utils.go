package main

import (
	"fmt"
	"math"
)

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
