package main

import (
	"fmt"
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
		s: &snake{
			length:     100.0,
			width:      20.0,
			speed:      0.5,
			color:      colornames.Limegreen,
			direction:  pixel.V(0.0, 10.0),
			constSpeed: 10.0,
		},
	}

	w.init("Snake 1.1")

	w.a = createApple(w.bounds)

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	last := time.Now()

	for !w.isEnded() {

		dt := time.Since(last).Seconds()
		last = time.Now()

		w.clear()

		w.processKeys(w.win)

		w.move(dt)

		w.draw(w.win)

		frames++
		select {
		case <-second:
			w.win.SetTitle(fmt.Sprintf("%s | FPS:%d | len:%d | speed:%v", w.cfg.Title, frames, len(w.s.pos), w.s.speed))
			frames = 0
		default:
		}
	}
}

func main() {
	pixelgl.Run(run)
}

/*

while() {

	if player.KeyPress() {
		snake.OnKeyPress(key)
	}


	if snake.HasEaten(apple) {
		snake.OnEaten(apple)
		createRandomApple()
	}

	snake.Move()

	clear()
	draw(game, 60)
}

*/
