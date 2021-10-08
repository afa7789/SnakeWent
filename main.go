package main

import (
	"flag"

	"github.com/afa7789/SnakeWent/pkg/game"
)

func main() {
	// Game Start
	width := flag.Int("width", 20, "width, horizontal size")
	height := flag.Int("height", 10, "height, vertical size")
	game.NewGame(*width, *height)
}
