package main

import (
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
    entityManager.SpawnEntity(gameWorld)

    player := entities.NewPlayer(
        gameWorld.(entities.IGameWorld),
        components.NewVelocityHandler())

    entityManager.SpawnEntity(player)
}

func update(entityManager entities.IEntityManager) {
    rl.BeginDrawing()

    rl.ClearBackground(constants.WindowBackgroundColor)

    for _, entity := range entityManager.GetEntities() {
        entity.Update()

        if renderable, ok := entity.(entities.IRenderable); ok {
            renderable.Render()
        }
    }

    rl.EndDrawing()
}
