package main

const (
	DT         = 1.0 / 60.0
	WINDOWSIZE = 256
	TILESIZE   = 16
)

const (
	GRAVITY    = 40
	JUMP_SPEED = -8
)

var TILETYPES = map[int]string{
	1: "Tile",
	2: "ToggleFloor",
	5: "Spikes",
	6: "Lever",
	9: "Player",
}

// var TEX *ebiten.Image
