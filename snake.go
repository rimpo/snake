package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
)

type snake struct {
	length        float64
	width         float64
	speed         float64
	color         color.RGBA
	direction     pixel.Vec
	prevDirection pixel.Vec
	pos           []pixel.Vec //start position,turn position, and end position
	constSpeed    float64
}

func createSnake() *snake {
	return &snake{
		length:     100.0,
		width:      20.0,
		speed:      0.5,
		color:      colornames.Limegreen,
		direction:  pixel.V(0.0, 10.0),
		constSpeed: 10.0,
	}
}

func (s *snake) initPos(pos ...pixel.Vec) {
	for _, p := range pos {
		s.pos = append(s.pos, p)
	}
}

func (s *snake) Draw(t pixel.Target) {
	imd := imdraw.New(nil)
	imd.EndShape = imdraw.RoundEndShape
	imd.Color = s.color

	var prev pixel.Vec
	for i, p := range s.pos {
		if i > 0 {
			imd.Push(prev, p)
		}
		prev = p
	}
	imd.Line(s.width)

	imd.Draw(t)
}

func (s *snake) OnKeyPress(btn pixelgl.Button) {
	switch btn {
	case pixelgl.KeyLeft:
		s.turnLeft()
	case pixelgl.KeyRight:
		s.turnRight()
	case pixelgl.KeyDown:
		s.turnDown()
	case pixelgl.KeyUp:
		s.turnUp()
	}
}

func (s *snake) IsCollidingWith(obj *Object) bool {
	return false
}

func (s *snake) OnCollision(obj *Object) {
}

func (s *snake) Move() {
}

func (s *snake) turnLeft() {
	if !(s.direction.X > 0.0) {
		s.direction = pixel.V(-s.speed, 0)
	}
}

func (s *snake) turnRight() {
	if !(s.direction.X < 0.0) {
		s.direction = pixel.V(s.speed, 0)
	}
}

func (s *snake) turnDown() {
	if !(s.direction.Y > 0.0) {
		s.direction = pixel.V(0, -s.speed)
	}
}

func (s *snake) turnUp() {
	if !(s.direction.Y < 0.0) {
		s.direction = pixel.V(0, s.speed)
	}
}

func (s *snake) head() pixel.Vec {
	return s.pos[0]
}

func (s *snake) moveTo(pos pixel.Vec) {
	s.pos = append([]pixel.Vec{pos}, s.pos...)
}

func (s *snake) move(delay float64) {
	//prepend a new co-ordinate
	if s.direction == s.prevDirection {
		s.pos[0] = s.pos[0].Add(s.direction.Scaled(delay))
	} else {
		s.pos = append([]pixel.Vec{s.pos[0].Add(s.direction.Scaled(delay))}, s.pos...)
	}

	var prev, p pixel.Vec
	total := 0.0
	i := 0

	//remove points outside the length of snake
	for i, p = range s.pos {
		if i > 0 {
			l := prev.Sub(p).Len()
			if total+l > s.length {
				diff := s.length - total
				//new tail according snake length
				if diff > 0 {
					diffVec := prev.To(p).Unit().Scaled(diff)
					s.pos[i] = prev.Add(diffVec)
				}
				break
			}
			total += l
		}
		prev = p
	}
	s.pos = s.pos[0 : i+1]
}
