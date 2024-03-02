package main

import (
	"fmt"

	"github.com/Ollie-Ave/Project_Kat/internal/components"
	"github.com/Ollie-Ave/Project_Kat/internal/constants"
	"github.com/Ollie-Ave/Project_Kat/internal/entities"
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
    setupWindow()
    
    entityManager := entities.NewEntityManager()
    initScene(entityManager)

    for !rl.WindowShouldClose() {
        update(entityManager)
        fmt.Println(rl.GetFPS())
    }

    rl.CloseWindow()
}

func setupWindow() {
    rl.InitWindow(constants.WindowWidth, constants.WindowHeight, constants.WindowTitle)

    rl.SetTargetFPS(constants.WindowTargetFPS)
    rl.SetExitKey(constants.WindowExitKey)
}

func initScene(entityManager entities.IEntityManager) {
    gameWorld := entities.NewGameWorld()
    entityManager.SpawnEntity("GameWorld", gameWorld)

    player := entities.NewPlayer(
        gameWorld.(entities.IGameWorld),
        components.NewVelocityHandler())

    entityManager.SpawnEntity("Player", player)

    entityManager.SpawnEntity("Camera", entities.NewCamera(entityManager))
}

func update(entityManager entities.IEntityManager) {
    rl.BeginDrawing()

    camera, err := entityManager.GetEntityById("Camera")

    if err != nil {
        panic("Camera not found")
    }

    camera.Update()
    rl.BeginMode2D(camera.(entities.ICamera).GetCamera())

    rl.ClearBackground(constants.WindowBackgroundColor)

    for _, entity := range entityManager.GetEntities() {
        entity.Update()

        if renderable, ok := entity.(entities.IRenderable); ok {
            renderable.Render()
        }
    }

    rl.EndMode2D()
    rl.EndDrawing()
}
