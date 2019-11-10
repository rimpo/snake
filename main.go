package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	rand.Seed(time.Now().Unix())

	w := &world{
		bounds:    pixel.R(10, 10, 1014, 758),
		winBounds: pixel.R(0, 0, 1024, 768),
		title:     "Snake 1.1",
	}

	w.a = createApple(w.bounds)
	w.s = createSnake()

	w.objs = append(w.objs, w.a)
	w.objs = append(w.objs, w.s)

	w.init()

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	last := time.Now()

	for !w.isEnded() {

		dt := time.Since(last).Seconds()
		last = time.Now()

		w.clear()

		w.processKeys()

		w.move(dt)

		w.draw()

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
