package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
	"image/color"
)

type wall struct {
	bounds pixel.Rect
	color  color.RGBA
	points []pixel.Vec
}

func createWall(bounds pixel.Rect) *wall {
	var points []pixel.Vec

	points = append(points, pixel.V(bounds.Min.X, bounds.Min.Y))
	points = append(points, pixel.V(bounds.Max.X, bounds.Min.Y))
	points = append(points, pixel.V(bounds.Max.X, bounds.Max.Y))
	points = append(points, pixel.V(bounds.Min.X, bounds.Max.Y))
	points = append(points, pixel.V(bounds.Min.X, bounds.Min.Y))

	return &wall{
		bounds: bounds,
		color:  colornames.Red,
		points: points,
	}
}

func (w *wall) Draw(t pixel.Target) {
	imd := imdraw.New(nil)
	imd.EndShape = imdraw.RoundEndShape
	imd.Color = w.color

	var prev pixel.Vec
	for i, p := range w.points {
		if i > 0 {
			imd.Push(prev, p)
		}
		prev = p
	}
	imd.Line(5.0)

	imd.Draw(t)
}

func (w *wall) IsInside(pt pixel.Vec) bool {
	//newHeadPos := w.s.pos[0].Add(w.s.direction)
	return w.bounds.Contains(pt)
}
