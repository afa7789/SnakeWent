package main

import (
	"github.com/afa7789/SnakeWent/pkg/game"
	"flag"
)

func main() {
	// Game Start
	width  := flag.Int("width", 10, "width, horizontal size")
	height := flag.Int("height", 10, "height, vertical size")
	game.NewGame(*width,*height)
}
