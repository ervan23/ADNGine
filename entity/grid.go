package entity

import rl "github.com/lachee/raylib-goplus/raylib"

type Grid struct {
	slices int
	spacing float32
}

func NewGrid(slices int, spacing float32) *Grid {
	return &Grid{slices, spacing}
}

func (g *Grid) Draw() {
	rl.DrawGrid(g.slices, g.spacing)
}

func (g *Grid) Update() {

}