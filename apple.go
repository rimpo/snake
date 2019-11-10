package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	//"image/color"
)

type apple struct {
	pos    pixel.Vec
	radius float64
	dead   bool
}

func (a *apple) Draw(t pixel.Target) {
	imd := imdraw.New(nil)
	imd.EndShape = imdraw.RoundEndShape
	imd.Color = colornames.Red
	imd.Push(a.pos)
	imd.Circle(a.radius, 0)
	imd.Draw(t)
}

func (a *apple) OnKeyPress(btn pixelgl.Button) {
	//Do nothing
}

func createApple(bounds pixel.Rect) *apple {
	return &apple{
		pos: pixel.V(
			float64(random(int(bounds.Min.X), int(bounds.Max.X))),
			float64(random(int(bounds.Min.Y), int(bounds.Max.Y))),
		),
		radius: 10.0,
		dead:   false,
	}
}

func (a *apple) isCollision(pos pixel.Vec) bool {
	if !a.dead && pos.To(a.pos).Len() <= a.radius {
		a.dead = true
		return true
	}
	return false
}
