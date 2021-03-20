package Render

import (
	of "OctaForceEngineGo"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"path/filepath"
	"runtime"
)

func Init() {
	//defer profile.Start().Stop()
	of.StartUp(start, update, stop, "Gravity Waves")
}

var absPath string
var camera int

func start() {
	_, b, _, _ := runtime.Caller(0)
	absPath = filepath.Dir(b)

	of.MaxFPS = 60
	of.MaxUPS = 30

	camera = of.CreateEntity()

	transform := of.Transform{}
	transform.Position = mgl32.Vec3{0, 0, 1500}
	transform.SetRotaion(mgl32.Vec3{0, 0, 0})
	of.AddComponent(camera, of.ComponentTransform, transform)
	of.AddComponent(camera, of.ComponentCamera, nil)
	of.SetActiveCameraEntity(camera)

	setUpPlane()
	setUpGravityObjects()
}

const (
	movementSpeed float32 = 100
	mouseSpeed    float32 = 3
)

var frame int
var lastKey = 0

const keyColldown = 10

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

	if of.KeyPressed(of.Key1) && (frame > lastKey+keyColldown) {
		cancelGravitation = !cancelGravitation
		lastKey = frame
	}

	if of.KeyPressed(of.Key2) {
		transform.Position = mgl32.Vec3{0, 0, 1500}
		transform.SetRotaion(mgl32.Vec3{0, 0, 0})
	} else if of.KeyPressed(of.Key3) {
		transform.Position = mgl32.Vec3{0, 300, 0}
		transform.SetRotaion(mgl32.Vec3{-90, 0, 0})
	}
	of.SetComponent(camera, of.ComponentTransform, transform)

	updateGravityObjects()
	updatePlane()

	frame++
}

func stop() {

}
