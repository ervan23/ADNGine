package entity

import (
	"adngine/color"
	rl "github.com/lachee/raylib-goplus/raylib"
)

type Rectangle struct {
	PosX   int
	PosY   int
	Width  int
	Height int
	Color  color.Color
	Lines  bool
}

func NewRectangle(posX, posY, width, height int, color color.Color, line bool) *Rectangle {
	return &Rectangle{posX, posY, width, height, color, line}
}

func (r *Rectangle) Draw() {
	if r.Lines {
		rl.DrawRectangleLines(r.PosX, r.PosY, r.Width, r.Height, rl.Color(r.Color))
	} else {
		rl.DrawRectangle(r.PosX, r.PosY, r.Width, r.Height, rl.Color(r.Color))
	}
}

func (r *Rectangle) Update() {

}
