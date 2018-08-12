package main

import (
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

type world struct {
	s         *snake
	a         *apple
	bounds    pixel.Rect
	winBounds pixel.Rect
}

func (w *world) init() {
	w.s = &snake{
		length: 100.0,
		width:  20.0,
		speed:  10.0,
		color:  colornames.Limegreen,
	}
	w.s.direction = pixel.V(0.0, w.s.speed)

	w.a = &apple{
		pos: pixel.V(
			float64(random(int(w.bounds.Min.X), int(w.bounds.Max.X))),
			float64(random(int(w.bounds.Min.Y), int(w.bounds.Max.Y))),
		),
		dead: false,
	}

	w.s.initPos(pixel.V(100, 100), pixel.V(200, 100))
}

func (w *world) draw(win pixel.Target) {
	w.s.draw(win)
	w.a.draw(win)
}

func (w *world) isWallCollision() bool {
	newHeadPos := w.s.pos[0].Add(w.s.direction)

	return !w.bounds.Contains(newHeadPos)
}

func (w *world) processKeys(win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeyLeft) {
		w.s.turnLeft()
	}
	if win.Pressed(pixelgl.KeyRight) {
		w.s.turnRight()
	}
	if win.Pressed(pixelgl.KeyDown) {
		w.s.turnDown()
	}
	if win.Pressed(pixelgl.KeyUp) {
		w.s.turnUp()
	}
	if !w.isWallCollision() {
		w.s.move()
	} else {
		//snake dead
		if w.s.pos[0].X < w.bounds.Min.X {
			w.s.pos[0].X = w.bounds.Min.X
		} else if w.s.pos[0].X > w.bounds.Max.X {
			w.s.pos[0].X = w.bounds.Max.X
		}
		if w.s.pos[0].Y < w.bounds.Min.Y {
			w.s.pos[0].Y = w.bounds.Min.Y
		} else if w.s.pos[0].Y > w.bounds.Max.Y {
			w.s.pos[0].Y = w.bounds.Max.Y
		}
	}

	if w.a.isCollision(w.s.pos[0]) {
		w.s.length += 50.0
		w.a.pos = pixel.V(
			float64(random(int(w.bounds.Min.X), int(w.bounds.Max.X))),
			float64(random(int(w.bounds.Min.Y), int(w.bounds.Max.Y))),
		)
		w.a.dead = false
	}
}
