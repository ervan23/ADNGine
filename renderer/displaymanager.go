package renderer

import (
	"adngine/color"
	rl "github.com/lachee/raylib-goplus/raylib"
	"log"
)

type DisplayManager struct {
	Width  int
	Height int
	Title  string
	FpsCap int
}

func NewDisplayManager(width, height, fpsCap int, title string) DisplayManager {
	log.Println("Init DisplayManager")
	dm := DisplayManager{width, height, title, fpsCap}
	dm.CreateDisplay()
	return dm
}

func (d *DisplayManager) CreateDisplay() {
	log.Println("Create Display ...")
	rl.InitWindow(d.Width, d.Height, d.Title)
	rl.SetTargetFPS(d.FpsCap)
}

func (d *DisplayManager) Open() bool {
	return rl.WindowShouldClose()
}

func (d *DisplayManager) Close() {
	rl.CloseWindow()
}

func (d *DisplayManager) ClearBackground(color color.Color) {
	rl.ClearBackground(rl.NewColor(color.R, color.G, color.B, color.A))
}

func (d *DisplayManager) Begin() {
	rl.BeginDrawing()
}

func (d *DisplayManager) End() {
	rl.EndDrawing()
}

func (d *DisplayManager) Begin3D(camera Camera3D) {
	rl.BeginMode3D(camera.Cam3D)
}

func (d *DisplayManager) End3D() {
	rl.EndMode3D()
}

func (d *DisplayManager) DrawFPS(posX, posY int) {
	rl.DrawFPS(posX, posY)
}