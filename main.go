package main

import (
	"adngine/color"
	"adngine/entity"
	"adngine/gameobject"
	"adngine/math"
	"adngine/renderer"
	"adngine/ui"
	rl "github.com/lachee/raylib-goplus/raylib"
)

type GameScene struct {
	camera        renderer.Camera3D
	cubePosition  math.Vector3
	mainBlock     entity.Entity
	lineBlock     entity.Entity
	mainGrid      entity.Entity
	mainRectangle entity.Entity
	lineRectangle entity.Entity
	texts         []entity.Entity
}

func (g *GameScene) Init() {
	g.camera = renderer.NewCamera3D(
		math.NewVector3(10.0, 10.0, 10.0),
		math.NewVector3(0.0, 0.0, 0.0),
		math.NewVector3(0.0, 1.0, 0.0),
		45.0,
		renderer.CAMERA_TYPE(renderer.CAMERA_PERSPECTIVE),
		renderer.CameraMode(renderer.CameraFree))

	g.cubePosition = math.NewVector3(0.0, 0.0, 0.0)
	g.mainBlock = gameobject.NewBlock(g.cubePosition, 2.0, 2.0, 2.0, color.Green, false)
	g.lineBlock = gameobject.NewBlock(g.cubePosition, 2.0, 2.0, 2.0, color.Red, true)
	g.mainGrid = gameobject.NewGrid(10, 1.0)
	g.mainRectangle = ui.NewRectangle(10, 10, 320, 133, color.Fade(color.SkyBlue, 0.5), false)
	g.lineRectangle = ui.NewRectangle(10, 10, 320, 133, color.Blue, true)

	g.texts = []entity.Entity{
		ui.NewText("Free camera default controls:", 20, 20, 10, color.Black),
		ui.NewText("- Mouse Wheel to Zoom in-out", 40, 40, 10, color.DarkGray),
		ui.NewText("- Mouse Wheel Pressed to Pan", 40, 60, 10, color.DarkGray),
		ui.NewText("- Alt + Mouse Wheel Pressed to Rotate", 40, 80, 10, color.DarkGray),
		ui.NewText("- Alt + Ctrl + Mouse Wheel Pressed for Smooth Zoom", 40, 100, 10, color.DarkGray),
		ui.NewText("- Z to zoom to (0, 0, 0)", 40, 120, 10, color.DarkGray),
	}
}

func (g *GameScene) Draw() {
	displayManager.Begin3D(g.camera)
	g.mainBlock.Draw()
	g.lineBlock.Draw()

	g.mainGrid.Draw()
	displayManager.End3D()

	g.mainRectangle.Draw()
	g.lineRectangle.Draw()

	for _, text := range g.texts {
		text.Draw()
	}

	displayManager.DrawFPS(10, 150)
}

func (g *GameScene) Update() {
	g.camera.UpdateCamera()
	if rl.IsKeyDown(rl.KeyX) {
		g.camera.Target = math.NewVector3(0, 0, 0)
	}
}

var displayManager = renderer.NewDisplayManager(800, 650, 10000, "Hello")

func main() {
	defer displayManager.Close()

	gameManager := renderer.NewGameManager()

	var gameScene renderer.Scene
	gameScene = &GameScene{}
	gameManager.Push(gameScene)

	for !displayManager.Open() {
		displayManager.Begin()
		displayManager.ClearBackground(color.RayWhite)
		gameManager.Draw()
		gameManager.Update()
		displayManager.End()
	}
}
