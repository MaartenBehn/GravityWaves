package main

import (
	of "OctaForceEngineGo"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"math/rand"
	"path/filepath"
	"runtime"
)

func main() {
	of.StartUp(start, update, stop)
}

var camera int

const (
	size  = 100
	scale = 10
)

var particles []int

func getParticle(x int, y int) int {
	return particles[x*size*2+y]
}
func setParticle(x int, y int, id int) {
	particles[x*size*2+y] = id
}

func start() {
	of.SetMaxFPS(60)
	of.SetMaxUPS(30)

	camera = of.CreateEntity()
	of.AddComponent(camera, of.ComponentCamera)
	transform := of.GetComponent(camera, of.ComponentTransform).(of.Transform)
	transform.SetPosition(mgl32.Vec3{0, 100, 500})
	of.SetComponent(camera, of.ComponentTransform, transform)
	of.SetActiveCameraEntity(camera)

	particles = make([]int, size*size*size*8)

	_, b, _, _ := runtime.Caller(0)
	absPath := filepath.Dir(b)
	mesh := of.LoadOBJ(absPath+"/mesh/Sphere.obj", false)

	particle := of.CreateEntity()

	of.AddComponent(particle, of.ComponentMesh)
	mesh.Material = of.Material{DiffuseColor: mgl32.Vec3{
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
	}}
	of.SetComponent(particle, of.ComponentMesh, mesh)

	particleTransform := of.GetComponent(particle, of.ComponentTransform).(of.Transform)
	particleTransform.SetPosition(
		mgl32.Vec3{float32(-size) * scale, 0, float32(-size) * scale})
	of.SetComponent(particle, of.ComponentTransform, particleTransform)

	setParticle(0, 0, particle)

	for x := 0; x < size*2; x++ {
		for y := 1; y < size*2; y++ {
			particle := of.CreateEntity()

			meshInstant := of.AddComponent(particle, of.ComponentMeshInstant).(of.MeshInstant)
			meshInstant.OwnEntity = particle
			meshInstant.MeshEntity = particles[0]
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

			setParticle(x, y, particle)
		}
	}
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
