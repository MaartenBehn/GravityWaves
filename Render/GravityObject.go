package Render

import (
	of "OctaForceEngineGo"
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

var gravityObjects []int
var gravityObjectColor mgl32.Vec3
var log [][]mgl32.Vec3
var logSize = 10

func setUpGravityObjects() {

	gravityObjects = make([]int, 2)
	gravityObjectColor = mgl32.Vec3{1, 0, 0}

	mesh := of.LoadOBJ(absPath+"/mesh/LowPolySphere.obj", false)
	for i, gravityObject := range gravityObjects {

		gravityObject = of.CreateEntity()

		of.AddComponent(gravityObject, of.ComponentMesh)
		mesh.Material = of.Material{DiffuseColor: gravityObjectColor}
		of.SetComponent(gravityObject, of.ComponentMesh, mesh)

		gravityObjects[i] = gravityObject
	}
}

var roundsPerFrame = 0.05
var radius = 10.0

func updateGravityObjects() {
	log = append(log, make([]mgl32.Vec3, 2))

	for i, gravityObject := range gravityObjects {
		transform := of.GetComponent(gravityObject, of.ComponentTransform).(of.Transform)

		transform.SetPosition(getPosOfGravityObject(frame, i))

		log[len(log)-1][i] = transform.GetPosition()
		of.SetComponent(gravityObject, of.ComponentTransform, transform)
	}

	for len(log) > logSize {
		log = append(log[:0], log[1:]...)
	}
}

func getPosOfGravityObject(frame int, nr int) mgl32.Vec3 {
	var pos mgl32.Vec3
	if nr == 0 {
		pos = mgl32.Vec3{
			float32(math.Sin(float64(frame)*roundsPerFrame) * radius),
			0,
			float32(math.Cos(float64(frame)*roundsPerFrame) * radius),
		}
	} else {
		pos = mgl32.Vec3{
			float32(-math.Sin(float64(frame)*roundsPerFrame) * radius),
			0,
			float32(-math.Cos(float64(frame)*roundsPerFrame) * radius),
		}
	}
	return pos
}
