package entities

import (
	"github.com/Ollie-Ave/Project_Kat/internal/constants"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameWorld struct {
    timePeriod bool
}

func NewGameWorld() IEntity {
    return &GameWorld {
        timePeriod: false,
    }
}

func (gameWorld *GameWorld) Update() {
}

func (gameWorld *GameWorld) Render() {
    floorHeight := gameWorld.GetFloorHeight(0)

    var worldColour rl.Color

    if gameWorld.timePeriod {
        worldColour = rl.Green
    } else {
        worldColour = rl.Orange
    }

    rl.DrawLine(0, int32(floorHeight), int32(rl.GetScreenWidth()), int32(floorHeight), worldColour)
}

func (gameWorld *GameWorld) ToggleTimePeriod() {
    gameWorld.timePeriod = !gameWorld.timePeriod
}

func (gameWorld *GameWorld) GetFloorHeight(xPosition float32) float32 {
    return float32(rl.GetScreenHeight() - constants.GameWorldFloorHeight)
}
