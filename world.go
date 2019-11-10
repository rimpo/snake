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
	cfg       pixelgl.WindowConfig
	win       *pixelgl.Window
}

func (w *world) init(title string) {
	w.cfg = pixelgl.WindowConfig{
		Title:  title,
		Bounds: w.winBounds,
		//VSync:  true,
	}
	var err error
	w.win, err = pixelgl.NewWindow(w.cfg)
	if err != nil {
		panic(err)
	}

	w.s.initPos(pixel.V(100, 100), pixel.V(200, 100))
}

func (w *world) clear() {
	w.win.Clear(colornames.Green)
}

func (w *world) draw(win pixel.Target) {
	w.s.draw(w.win)
	w.a.draw(w.win)

	w.win.Update()
}

func (w *world) isEnded() bool {
	return w.win.Closed()
}

func (w *world) isWallCollision() bool {
	newHeadPos := w.s.pos[0].Add(w.s.direction)

	return !w.bounds.Contains(newHeadPos)
}

func (w *world) processKeys(win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyLeft) {
		w.s.turnLeft()
	}
	if win.JustPressed(pixelgl.KeyRight) {
		w.s.turnRight()
	}
	if win.JustPressed(pixelgl.KeyDown) {
		w.s.turnDown()
	}
	if win.JustPressed(pixelgl.KeyUp) {
		w.s.turnUp()
	}

}

func (w *world) move(delay float64) {
	//w.s.constSpeed = delay * 500
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
		w.s.constSpeed += w.s.constSpeed * 0.1
		w.s.speed += w.s.constSpeed * delay

		w.a.pos = pixel.V(
			float64(random(int(w.bounds.Min.X), int(w.bounds.Max.X))),
			float64(random(int(w.bounds.Min.Y), int(w.bounds.Max.Y))),
		)
		w.a.dead = false
	}
	w.s.prevDirection = w.s.direction
}
