package camera

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func CameraInit(r *RayRectangle) {
  player := rl.Rectangle{ 400, 280, 40, 40 }
  var buildings [MAX_BUILDINGS]rl.Rectangle

  var buildColor [MAX_BUILDINGS]rl.Color 
  spacing := 0

  for i := 0; i < MAX_BUILDINGS; i++ {
    buildings[i].Width = float32(rl.GetRandomValue(50,200))
    buildings[i].Height = float32(rl.GetRandomValue(100, 800))
    buildings[i].Y = float32(r.ScreenHeight) - 130.0 - buildings[i].Height
    buildings[i].X = float32(-6000.0) + float32(spacing)
 
    spacing += int(buildings[i].Width)
    
    buildColor[i] = rl.Color{ uint8(rl.GetRandomValue(200,240)), uint8(rl.GetRandomValue(200,240)), uint8(rl.GetRandomValue(200,250)), 255 }
  }

  camera := rl.Camera2D{ 0 }
  
  camera.Target = rl.Vector2({ player.X + 20.0 player.Y + 20.0 })
}
