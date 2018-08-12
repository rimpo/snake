package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
	//"image/color"
)

type apple struct {
	pos  pixel.Vec
	dead bool
}

func (a *apple) draw(t pixel.Target) {
	imd := imdraw.New(nil)
	imd.EndShape = imdraw.RoundEndShape
	imd.Color = colornames.Red
	imd.Push(a.pos)
	imd.Circle(20.0, 0)
	imd.Draw(t)
}

func (a *apple) isCollision(pos pixel.Vec) bool {
	if !a.dead && pos.To(a.pos).Len() <= 20.0 {
		a.dead = true
		return true
	}
	return false
}
