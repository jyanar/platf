package main

const (
	DT          = 1.0 / 60.0
	WINDOWSIZE  = 256
	WINDOWSCALE = 3
	TILESIZE    = 16
)

const (
	GRAVITY        = 8
	JUMP_SPEED     = -12
	MAX_SPEED      = 20
	ACCELERATION   = 100
	MAX_FALL_SPEED = 25
	FRICTION       = 6
)

var TILETYPES = map[int]string{
	1: "Tile",
	2: "ToggleFloor",
	5: "Spikes",
	6: "Lever",
	9: "Player",
}
