package Render

import (
	of "OctaForceEngineGo"
	"github.com/go-gl/mathgl/mgl32"
)

var plane []int
var baseColor mgl32.Vec3

const (
	size  = 30
	scale = 10
)

func getPointInPlane(x int, y int) int {
	return plane[x*size*2+y]
}
func setPointInPlane(x int, y int, id int) {
	plane[x*size*2+y] = id
}
func setUpPlane() {
	plane = make([]int, size*size*4)
	baseColor = mgl32.Vec3{0.5, 0.5, 0.5}

	point := of.CreateEntity()
	mesh := of.LoadOBJ(absPath+"/mesh/LowPolySphere.obj", false)
	of.AddComponent(point, of.ComponentMesh)
	mesh.Material = of.Material{DiffuseColor: baseColor}
	of.SetComponent(point, of.ComponentMesh, mesh)

	pointTransform := of.GetComponent(point, of.ComponentTransform).(of.Transform)
	pointTransform.SetPosition(
		mgl32.Vec3{float32(-size) * scale, 0, float32(-size) * scale})
	of.SetComponent(point, of.ComponentTransform, pointTransform)

	setPointInPlane(0, 0, point)

	for x := 0; x < size*2; x++ {
		for y := 0; y < size*2; y++ {
			if x == 0 && y == 0 {
				continue
			}
			particle := of.CreateEntity()

			meshInstant := of.AddComponent(particle, of.ComponentMeshInstant).(of.MeshInstant)
			meshInstant.OwnEntity = particle
			meshInstant.MeshEntity = plane[0]
			meshInstant.Material = of.Material{DiffuseColor: baseColor}
			of.SetComponent(particle, of.ComponentMeshInstant, meshInstant)

			particleTransform := of.GetComponent(particle, of.ComponentTransform).(of.Transform)
			particleTransform.SetPosition(
				mgl32.Vec3{float32(x-size) * scale, 0, float32(y-size) * scale})
			of.SetComponent(particle, of.ComponentTransform, particleTransform)

			setPointInPlane(x, y, particle)
		}
	}
}

var gravityMultiplier float32 = 1000.0
var cancelGravitation = true

func updatePlane() {

	for _, point := range plane {
		transform := of.GetComponent(point, of.ComponentTransform).(of.Transform)
		pos := transform.GetPosition()
		pos[1] = 0

		if cancelGravitation {
			distance := mgl32.Vec3{0, 0, 0}.Sub(mgl32.Vec3{pos[0], 0, pos[2]}).Len()
			pos[1] += 2.0 / distance * gravityMultiplier
		}

		for _, gravityObject := range gravityObjects {

			gravityObjectTransform := of.GetComponent(gravityObject, of.ComponentTransform).(of.Transform)
			distance := gravityObjectTransform.GetPosition().Sub(mgl32.Vec3{pos[0], 0, pos[2]}).Len()
			pos[1] -= 1.0 / distance * gravityMultiplier

		}
		transform.SetPosition(pos)
		of.SetComponent(point, of.ComponentTransform, transform)
	}

}
