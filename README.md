# platf

This is an experiment in building a simple platformer using the Go programming language and the Ebiten engine ([ebitengine.org](https://ebitengine.org)). It roughly follows the [LOVE Platformer Guide](https://0x72.itch.io/love2d-platformer-guide) by 0x72.

The project features a simple state manager (for switching between pause, play, and dead states), entities that can be interacted with (such as spikes, levers, and floors), different levels, and enemies that move. Many things could certainly be done better, but this is my first project in Go.

But, I've had a lot of fun! Go is simple, and because of that, feels quite elegant to use. Relative to building something similar using Lua/Love2D, I appreciate Go's static typing, robust standard library, and standardized styling across codebases via `gofmt`.
