package math

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

func NewVector3(X, Y, Z float32) Vector3 {
	return Vector3{X, Y, Z}
}