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
		winBounds: pixel.R(0, 0, 1024, 768),
		title:     "Snake 1.1",
	}

	w.s = createSnake()
	w.wall = createWall(pixel.R(50, 50, 974, 718))
	w.a = createApple(w.wall.bounds)

	//moveable objects
	w.moveableObjs = append(w.moveableObjs, w.s)

	//non-moveable objects
	w.objects = append(w.objects, w.a)
	w.objects = append(w.objects, w.wall)

	w.init()

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	last := time.Now()

	for !w.isEnded() {
		//calculate delta to control FPS1000ms/60
		dt := time.Since(last)
		last = time.Now()
		time.Sleep(time.Duration(16666*time.Microsecond) - dt)
		delta := dt.Seconds() + 1.0

		w.clear()

		w.processKeys()

		w.move(delta)

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
