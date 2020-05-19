package entity

import (
	"adngine/color"
	"adngine/math"
	rl "github.com/lachee/raylib-goplus/raylib"
)

type Block struct {
	Position math.Vector3
	Width    float32
	Height   float32
	Length   float32
	Color    color.Color
	Wires	 bool
}

func NewBlock(position math.Vector3, width, height, length float32, color color.Color, wires bool) *Block {
	return &Block{position, width, height, length, color, wires}
}

func (b *Block) Draw() {
	if b.Wires {
		rl.DrawCubeWires(rl.Vector3(b.Position), b.Width, b.Height, b.Length, rl.Color(b.Color))
	} else {
		rl.DrawCube(rl.Vector3(b.Position), b.Width, b.Height, b.Length, rl.Color(b.Color))
	}
}

func (b *Block) Update() {

}