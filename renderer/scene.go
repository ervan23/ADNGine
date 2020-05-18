package renderer

type Scene interface {
	Init()
	Draw()
	Update()
}