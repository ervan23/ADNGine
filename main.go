package main

import (
	"adngine/color"
	"adngine/renderer"
)

// EXAMPLE OF MAIN [Not work at this stade]

func main() {
	displayManager := renderer.NewDisplayManager(800, 650, 60, "Hello")
	defer displayManager.Close()

	gameManager := renderer.NewGameManager()

	for !displayManager.Open() {
		displayManager.Begin()
		displayManager.ClearBackground(color.RayWhite)
		gameManager.Draw()
		gameManager.Update()
		displayManager.End()
	}
}
