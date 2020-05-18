package main

import (
	"adngine/color"
	"adngine/entity"
	"adngine/math"
	"adngine/renderer"
)

type GameScene struct {
	camera       renderer.Camera3D
	cubePosition math.Vector3
	mainBlock 	 entity.Entity
}

func (g *GameScene) Init() {
	g.camera = renderer.NewCamera3D(
		math.NewVector3(10.0, 10.0, 10.0),
		math.NewVector3(0.0, 0.0, 0.0),
		math.NewVector3(0.0, 1.0, 0.0),
		45.0,
		renderer.CAMERA_TYPE(renderer.CAMERA_PERSPECTIVE),
		renderer.CameraMode(renderer.CameraFirstPerson))

	g.cubePosition = math.NewVector3(0.0, 0.0, 0.0)
	g.mainBlock  = entity.NewBlock(g.cubePosition, 2.0, 2.0, 2.0, color.Red)
}

func (g *GameScene) Draw() {
	displayManager.Begin3D(g.camera)
	g.mainBlock.Draw()
	displayManager.End3D()}

func (g *GameScene) Update() {
	g.camera.UpdateCamera()
}

var displayManager = renderer.NewDisplayManager(800, 650, 60, "Hello")

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
		displayManager.End3D()
		displayManager.End()
	}
}
