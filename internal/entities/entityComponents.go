package entities

import rl "github.com/gen2brain/raylib-go/raylib"

type IEntity interface {
    Update()
}

type IRenderable interface {
    Render()
}

type IGameWorld interface {
    GetFloorHeight(xPosition float32) float32

    ToggleTimePeriod()
}

type ICollidable interface {
    GetHitboxCenter() rl.Vector2
}
