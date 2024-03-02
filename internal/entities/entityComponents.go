package entities

type IEntity interface {
    Update()
}

type IRenderable interface {
    Render()
}

type IGameWorld interface {
    GetFloorHeight(xPosition float32) float32
}
