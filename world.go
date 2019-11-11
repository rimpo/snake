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
	IsCollidingWith(Object, float64) bool
	OnCollision(Object)
}

type MoveableObject interface {
	Object
	CollisionObject
	Move(float64)
	OnKeyPress(pixelgl.Button)
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

	w.s.initPos(pixel.V(100, 100), pixel.V(200, 100), pixel.V(200, 200))
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

func (w *world) move(delta float64) {

	//check and handle if moveable object collide
	for i, _ := range w.moveableObjs {
		for j, _ := range w.objects {
			if w.moveableObjs[i].IsCollidingWith(w.objects[j], delta) {
				w.moveableObjs[i].OnCollision(w.objects[j])
			}
		}
		w.moveableObjs[i].Move(delta)
	}

}
