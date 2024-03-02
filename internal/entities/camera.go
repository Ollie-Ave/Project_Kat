package entities

import (
	"github.com/Ollie-Ave/Project_Kat/internal/constants"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ICamera interface {
    GetCamera() rl.Camera2D
}

type Camera struct {
    camera rl.Camera2D
    entityManager IEntityManager
}

func NewCamera(entityManager IEntityManager) IEntity {
    player, err := entityManager.GetEntityById("Player")

    if err != nil {
        panic("Player not found")
    }

    return &Camera {
        camera: rl.NewCamera2D(
            rl.NewVector2((constants.WindowWidth / 2) - (constants.PlayerHitboxWidth / 2), (constants.WindowHeight / 2) + constants.CameraYOffset),
            player.(ICollidable).GetHitboxCenter(),
            0,
            1),
        entityManager: entityManager,
    }
}

func (camera *Camera) GetCamera() rl.Camera2D {
    return camera.camera
}

func (camera *Camera) Update() {
    player, err := camera.entityManager.GetEntityById("Player")

    if err != nil {
        panic("Player not found")
    }

    camera.camera.Target = player.(ICollidable).GetHitboxCenter()
}
