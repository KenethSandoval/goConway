package main

import rl "github.com/gen2brain/raylib-go/raylib"

type RayWindow struct {
  ScreenWidth     int32
  ScreenHeight    int32
}

func main() {
  rayWindow := RayWindow{ ScreenWidth: 800, ScreenHeight: 400 }

  rl.InitWindow(rayWindow.ScreenWidth, rayWindow.ScreenHeight, "raylib [core] test")

  ballPosition := rl.Vector2{ X: float32(rayWindow.ScreenWidth) /2, Y: float32(rayWindow.ScreenHeight) /2 }

  rl.SetTargetFPS(60)

  for !rl.WindowShouldClose() {
    if rl.IsKeyDown(rl.KeyRight) {
      ballPosition.X += 2.0
    }

    if rl.IsKeyDown(rl.KeyLeft) {
      ballPosition.X -= 2.0
    }

    if rl.IsKeyDown(rl.KeyUp) {
      ballPosition.Y -= 2.0
    }

    if rl.IsKeyDown(rl.KeyDown) {
      ballPosition.Y += 2.0
    }

    rl.BeginDrawing()
    rl.ClearBackground(rl.RayWhite)
    rl.DrawText("Mueve la pelota con las teclas de flecha", 10, 10, 20, rl.DarkGray)
    rl.DrawCircleV(ballPosition, 50, rl.Maroon)

    rl.EndDrawing()
  }

  rl.CloseWindow()
}

