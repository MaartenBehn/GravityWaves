package main

import (
	of "OctaForceEngineGo"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/pkg/profile"
	"math/rand"
	"path/filepath"
	"runtime"
)

func main() {
	defer profile.Start().Stop()
	of.StartUp(start, update, stop)
}

var absPath string
var camera int

const (
	size  = 100
	scale = 10
)

var plane []int

func getPointInPlane(x int, y int) int {
	return plane[x*size*2+y]
}
func setPointInPlane(x int, y int, id int) {
	plane[x*size*2+y] = id
}
func setUpPlane() {
	plane = make([]int, size*size*size*8)

	point := of.CreateEntity()

	mesh := of.LoadOBJ(absPath+"/mesh/LowPolySphere.obj", false)
	of.AddComponent(point, of.ComponentMesh)
	mesh.Material = of.Material{DiffuseColor: mgl32.Vec3{
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
	}}
	of.SetComponent(point, of.ComponentMesh, mesh)

	pointTransform := of.GetComponent(point, of.ComponentTransform).(of.Transform)
	pointTransform.SetPosition(
		mgl32.Vec3{float32(-size) * scale, 0, float32(-size) * scale})
	of.SetComponent(point, of.ComponentTransform, pointTransform)

	setPointInPlane(0, 0, point)

	for x := 0; x < size*2; x++ {
		for y := 1; y < size*2; y++ {
			particle := of.CreateEntity()

			meshInstant := of.AddComponent(particle, of.ComponentMeshInstant).(of.MeshInstant)
			meshInstant.OwnEntity = particle
			meshInstant.MeshEntity = plane[0]
			meshInstant.Material = of.Material{DiffuseColor: mgl32.Vec3{
				rand.Float32(),
				rand.Float32(),
				rand.Float32(),
			}}
			of.SetComponent(particle, of.ComponentMeshInstant, meshInstant)

			particleTransform := of.GetComponent(particle, of.ComponentTransform).(of.Transform)
			particleTransform.SetPosition(
				mgl32.Vec3{float32(x-size) * scale, 0, float32(y-size) * scale})
			of.SetComponent(particle, of.ComponentTransform, particleTransform)

			setPointInPlane(x, y, particle)
		}
	}
}

func start() {
	_, b, _, _ := runtime.Caller(0)
	absPath = filepath.Dir(b)

	of.SetMaxFPS(60)
	of.SetMaxUPS(30)

	camera = of.CreateEntity()
	of.AddComponent(camera, of.ComponentCamera)
	transform := of.GetComponent(camera, of.ComponentTransform).(of.Transform)
	transform.SetPosition(mgl32.Vec3{0, 100, 500})
	of.SetComponent(camera, of.ComponentTransform, transform)
	of.SetActiveCameraEntity(camera)

	setUpPlane()
}

const (
	movementSpeed float32 = 100
	mouseSpeed    float32 = 3
)

func update() {
	fmt.Printf("FPS: %f UPS: %f \r", of.GetFPS(), of.GetUPS())

	deltaTime := float32(of.GetDeltaTime())

	transform := of.GetComponent(camera, of.ComponentTransform).(of.Transform)
	if of.KeyPressed(of.KeyW) {
		transform.MoveRelative(mgl32.Vec3{0, 0, -1}.Mul(deltaTime * movementSpeed))
	}
	if of.KeyPressed(of.KeyS) {
		transform.MoveRelative(mgl32.Vec3{0, 0, 1}.Mul(deltaTime * movementSpeed))
	}
	if of.KeyPressed(of.KeyA) {
		transform.MoveRelative(mgl32.Vec3{-1, 0, 0}.Mul(deltaTime * movementSpeed))
	}
	if of.KeyPressed(of.KeyD) {
		transform.MoveRelative(mgl32.Vec3{1, 0, 0}.Mul(deltaTime * movementSpeed))
	}
	if of.MouseButtonPressed(of.MouseButtonLeft) {
		mouseMovement := of.GetMouseMovement()
		transform.Rotate(mgl32.Vec3{-1, 0, 0}.Mul(mouseMovement.Y() * deltaTime * mouseSpeed))
		transform.Rotate(mgl32.Vec3{0, -1, 0}.Mul(mouseMovement.X() * deltaTime * mouseSpeed))
	}
	of.SetComponent(camera, of.ComponentTransform, transform)
}

func stop() {

}
