package Render

import (
	of "OctaForceEngineGo"
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

var gravityObjects []int
var gravityObjectColor mgl32.Vec3

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

var roundsPerFrame = 0.1
var frame int

func updateGravityObjects() {
	for i, gravityObject := range gravityObjects {
		transform := of.GetComponent(gravityObject, of.ComponentTransform).(of.Transform)
		if i == 0 {
			transform.SetPosition(mgl32.Vec3{
				float32(math.Sin(float64(frame)*roundsPerFrame) * 10),
				0,
				float32(math.Cos(float64(frame)*roundsPerFrame) * 10),
			})
		} else {
			transform.SetPosition(mgl32.Vec3{
				float32(-math.Sin(float64(frame)*roundsPerFrame) * 10),
				0,
				float32(-math.Cos(float64(frame)*roundsPerFrame) * 10),
			})
		}

		of.SetComponent(gravityObject, of.ComponentTransform, transform)
	}
	frame++
}
