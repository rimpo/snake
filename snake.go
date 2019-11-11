package main

import (
	"fmt"
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
	dead          bool
}

func createSnake() *snake {
	return &snake{
		length:     100.0,
		width:      5.0,
		speed:      1.0,
		color:      colornames.Limegreen,
		direction:  pixel.V(0.0, 10.0),
		constSpeed: 10.0,
		dead:       false,
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

func (s *snake) IsCollidingWith(obj Object, delta float64) bool {
	switch o := obj.(type) {
	case *wall:
		if !o.IsInside(s.head().Scaled(delta)) {
			//fmt.Println("wall collision:", s.head(), delta)
			return true
		}
	case *apple:
		if o.IsInside(s.head().Scaled(delta)) {
			fmt.Println(s.head().Scaled(delta))
			return true
		}
	}
	return false
}

func (s *snake) OnCollision(obj Object) {
	switch o := obj.(type) {
	case *wall:
		if s.pos[0].X < o.bounds.Min.X {
			s.pos[0].X = o.bounds.Min.X
		} else if s.pos[0].X > o.bounds.Max.X {
			s.pos[0].X = o.bounds.Max.X
		}
		if s.pos[0].Y < o.bounds.Min.Y {
			s.pos[0].Y = o.bounds.Min.Y
		} else if s.pos[0].Y > o.bounds.Max.Y {
			s.pos[0].Y = o.bounds.Max.Y
		}
		s.dead = true
	case *apple:
		s.length += 50.0
		s.constSpeed += s.constSpeed * 0.1
		//w.s.speed += w.s.constSpeed * delay

		/*w.a.pos = pixel.V(
			float64(random(int(w.wall.bounds.Min.X), int(w.wall.bounds.Max.X))),
			float64(random(int(w.wall.bounds.Min.Y), int(w.wall.bounds.Max.Y))),
		)
		o.dead = false
		*/

	}
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

func (s *snake) Move(delta float64) {
	if s.dead {
		return
	}
	//prepend a new co-ordinate
	if s.direction == s.prevDirection {
		s.pos[0] = s.pos[0].Add(s.direction.Scaled(delta))
	} else {
		s.pos = append([]pixel.Vec{s.pos[0].Add(s.direction.Scaled(delta))}, s.pos...)
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
	s.prevDirection = s.direction
}
