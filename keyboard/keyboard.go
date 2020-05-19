package keyboard

import rl "github.com/lachee/raylib-goplus/raylib"

func IsKeyUp(key Key) bool {
	return rl.IsKeyUp(rl.Key(key))
}

func IsKeyDown(key Key) bool {
	return rl.IsKeyDown(rl.Key(key))
}

func IsKeyPressed(key Key) bool {
	return rl.IsKeyPressed(rl.Key(key))
}

func IsKeyReleased(key Key) bool {
	return rl.IsKeyReleased(rl.Key(key))
}