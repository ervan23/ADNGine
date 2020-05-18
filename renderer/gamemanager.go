package renderer

import "log"

type Stack []Scene

func (s Stack) Peek() Scene {
	return s[len(s)-1]
}

func (s *Stack) Pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *Stack) Push(scene Scene) {
	*s = append(*s, scene)
}

type GameManager struct {
	scenes Stack
}

func NewGameManager() GameManager {
	return GameManager{Stack{}}
}

func (g *GameManager) Push(s Scene) {
	s.Init()
	log.Println("Init Scene")
	g.scenes.Push(s)
}

func (g *GameManager) Pop()  {
	g.scenes.Pop()
}

func (g *GameManager) Set(s Scene) {
	g.scenes.Pop()
	s.Init()
	g.scenes.Push(s)
}

func (g *GameManager) Draw() {
	scene := g.scenes.Peek()
	scene.Draw()
}

func (g *GameManager) Update() {
	scene := g.scenes.Peek()
	scene.Update()
}
