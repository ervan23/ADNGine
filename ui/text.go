package ui

import (
	"adngine/color"
	rl "github.com/lachee/raylib-goplus/raylib"
)

type Text struct {
	Title    string
	PosX    int
	PosY   int
	FontSize int
	Color    color.Color
}

func NewText(title string, posX, posY, fontSize int, color color.Color) *Text {
	return &Text{title, posX, posY, fontSize, color}
}

func (t *Text) Draw() {
	rl.DrawText(t.Title, t.PosX, t.PosY, t.FontSize, rl.Color(t.Color))
}

func (t *Text) Update() {

}