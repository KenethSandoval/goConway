package main

import (
	"github.com/KenethSandoval/goConway/camera"
	"github.com/KenethSandoval/goConway/window"
)


func main() {
  rayWindow := window.RayWindow{ ScreenWidth: 800, ScreenHeight: 400 }
  window.CreateWindow(&rayWindow)

  camera.CameraInit()
}

