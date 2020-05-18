package renderer

import (
	"adngine/math"
	rl "github.com/lachee/raylib-goplus/raylib"
)

type CAMERA_TYPE int32

const (
	CAMERA_PERSPECTIVE  = rl.CameraTypePerspective
	CAMERA_ORTHOGRAPHIC = rl.CameraTypeOrthographic
)

type CameraMode int32

const (
	CameraFree        = rl.CameraFree
	CameraOrbital     = rl.CameraOrbital
	CameraFirstPerson = rl.CameraFirstPerson
	CameraThirdPerson = rl.CameraThirdPerson
)

type Camera3D struct {
	Position math.Vector3
	Target   math.Vector3
	Up       math.Vector3
	Fovy     float32
	Type     CAMERA_TYPE
	Cam3D 	 rl.Camera
}

func NewCamera3D(position, target, up math.Vector3, fovy float32, cameraType CAMERA_TYPE, cameraMode CameraMode) Camera3D {
	cam := Camera3D{position, target, up, fovy, cameraType, rl.Camera{}}
	cam.InitCamera(cameraMode)
	return cam
}

func (c *Camera3D) InitCamera(cameraMode CameraMode) {
	cam3D := rl.NewCamera(rl.Vector3(c.Position), rl.Vector3(c.Target), rl.Vector3(c.Up), c.Fovy, rl.CameraType(c.Type))
	rl.SetCameraMode(&cam3D, rl.CameraMode(cameraMode))
	c.Cam3D = cam3D
}

func (c *Camera3D) UpdateCamera() {
	rl.UpdateCamera(&c.Cam3D)
}
