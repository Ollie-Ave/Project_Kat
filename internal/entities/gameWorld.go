package entities

import (
	"github.com/Ollie-Ave/Project_Kat/internal/constants"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameWorld struct {
}

func NewGameWorld() IEntity {
    return &GameWorld {}
}

func (gameWorld *GameWorld) Update() {
}

func (gameWorld *GameWorld) Render() {
    floorHeight := gameWorld.GetFloorHeight(0)

    rl.DrawLine(0, int32(floorHeight), int32(rl.GetScreenWidth()), int32(floorHeight), rl.Green)
}

func (gameWorld *GameWorld) GetFloorHeight(xPosition float32) float32 {
    return float32(rl.GetScreenHeight() - constants.GameWorldFloorHeight)
}
