package main

import (
	//	"fmt"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	rand.Seed(time.Now().Unix())

	w := &world{
		bounds:    pixel.R(10, 10, 1014, 758),
		winBounds: pixel.R(0, 0, 1024, 768),
	}

	w.init()

	cfg := pixelgl.WindowConfig{
		Title:  "Snake 1.0",
		Bounds: w.winBounds,
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Clear(colornames.Green)

		w.processKeys(win)

		w.draw(win)
		win.Update()

		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	pixelgl.Run(run)
}
