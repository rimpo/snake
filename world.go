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

type Object interface {
	Draw(t pixel.Target)
}

type CollisionObject interface {
	IsCollidingWith(*Object) bool
	OnCollision(*Object)
}

type MoveableObject interface {
	Object
	CollisionObject
	Move()
	OnKeyPress(btn pixelgl.Button)
}

type world struct {
	s            *snake
	a            *apple
	wall         *wall
	winBounds    pixel.Rect
	cfg          pixelgl.WindowConfig
	win          *pixelgl.Window
	title        string
	objects      []Object
	moveableObjs []MoveableObject
}

func (w *world) init() {
	w.cfg = pixelgl.WindowConfig{
		Title:  w.title,
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

func (w *world) draw() {
	for i, _ := range w.objects {
		w.objects[i].Draw(w.win)
	}

	for i, _ := range w.moveableObjs {
		w.moveableObjs[i].Draw(w.win)
	}
	w.win.Update()
}

func (w *world) isEnded() bool {
	return w.win.Closed()
}

func (w *world) isWallCollision() bool {
	newHeadPos := w.s.pos[0].Add(w.s.direction)

	return !w.wall.bounds.Contains(newHeadPos)
}

func (w *world) processKeys() {
	var btn pixelgl.Button
	if w.win.JustPressed(pixelgl.KeyLeft) {
		btn = pixelgl.KeyLeft
	} else if w.win.JustPressed(pixelgl.KeyRight) {
		btn = pixelgl.KeyRight
	} else if w.win.JustPressed(pixelgl.KeyDown) {
		btn = pixelgl.KeyDown
	} else if w.win.JustPressed(pixelgl.KeyUp) {
		btn = pixelgl.KeyUp
	}
	for i, _ := range w.moveableObjs {
		w.moveableObjs[i].OnKeyPress(btn)
	}
}

func (w *world) move(delay float64) {
	/*
		for i, _ := range w.moveableObjs {
			w.moveableObjs[i].Move()
		}

		//handle collision
		for i, _ := range w.moveableObjs {
				for j, _ := range w.objects {
					if w.objects[j].IsInside(w.moveableObjs[i]) {
						w.moveableObjs[i].OnCollision(w.objects[j])
					}
				}

		}

	*/
	//w.s.constSpeed = delay * 500
	if !w.isWallCollision() {
		w.s.move(delay)
	} else {
		//snake dead
		if w.s.pos[0].X < w.wall.bounds.Min.X {
			w.s.pos[0].X = w.wall.bounds.Min.X
		} else if w.s.pos[0].X > w.wall.bounds.Max.X {
			w.s.pos[0].X = w.wall.bounds.Max.X
		}
		if w.s.pos[0].Y < w.wall.bounds.Min.Y {
			w.s.pos[0].Y = w.wall.bounds.Min.Y
		} else if w.s.pos[0].Y > w.wall.bounds.Max.Y {
			w.s.pos[0].Y = w.wall.bounds.Max.Y
		}
		w.s.move(delay)
	}

	if w.a.IsInside(w.s.pos[0]) {
		w.s.length += 50.0
		w.s.constSpeed += w.s.constSpeed * 0.1
		//w.s.speed += w.s.constSpeed * delay

		w.a.pos = pixel.V(
			float64(random(int(w.wall.bounds.Min.X), int(w.wall.bounds.Max.X))),
			float64(random(int(w.wall.bounds.Min.Y), int(w.wall.bounds.Max.Y))),
		)
		w.a.dead = false
	}
	w.s.prevDirection = w.s.direction
}
