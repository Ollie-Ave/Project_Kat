package entities

import (
	"github.com/Ollie-Ave/Project_Kat/internal/components"
	"github.com/Ollie-Ave/Project_Kat/internal/constants"
	"github.com/Ollie-Ave/Project_Kat/internal/shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
    gameWorld IGameWorld

    hitbox rl.Rectangle
    velocityHandler components.IVelocityHandler
}

func NewPlayer(gameWorld IGameWorld, velocityHandler components.IVelocityHandler) IEntity {
    velocityHandler.CreateVelocityType(
        "walk",
        components.NewVelocity(rl.NewVector2(constants.PlayerMaxSpeed, 0),
        constants.PlayerFriction))

    velocityHandler.CreateVelocityType(
        "gravity",
        components.NewVelocity(rl.NewVector2(0, constants.PlayerTerminalFallingVelocity),
        constants.PlayerFriction))

    return &Player {
        hitbox: rl.NewRectangle(100, 200, constants.PlayerHitboxWidth, constants.PlayerHitboxHeight),
        velocityHandler: velocityHandler,

        gameWorld: gameWorld,
    }
}

func (player *Player) GetHitboxCenter() rl.Vector2 {
    return rl.NewVector2(
        player.hitbox.X + player.hitbox.Width/2,
        player.hitbox.Y + player.hitbox.Height/2)
}

func (player *Player) Update() {
    player.velocityHandler.HandleFriction()

    player.updatePhysics()

    player.updateHitboxPosition()
}

func (player *Player) Render() {
    rl.DrawRectangleLines(
        int32(player.hitbox.X),
        int32(player.hitbox.Y),
        int32(player.hitbox.Width),
        int32(player.hitbox.Height),
        rl.Orange)
}

func (player *Player) updateHitboxPosition() {
    velocity := player.velocityHandler.GetCombinedVelocity()

    player.hitbox.X += velocity.X
    
    player.hitbox.Y = shared.ClampValue(
        velocity.Y + player.hitbox.Y,
        0,
        player.gameWorld.GetFloorHeight(player.hitbox.X) - player.hitbox.Height)
}

func (player *Player) updatePhysics() {
    player.handleInput()
    player.handleGravity()
}

func (player *Player) handleInput() {
    player.handleWalk()
    player.handleJump()
    player.handleAbilities()

}
func (player *Player) handleAbilities() {
    if rl.IsKeyPressed(rl.KeyR) {
        player.gameWorld.ToggleTimePeriod()
    }
}

func (player *Player) handleJump() {
    if rl.IsKeyPressed(rl.KeySpace) && player.isTouchingGround() {
        player.velocityHandler.SetVelocity(
            "gravity",
            rl.NewVector2(0, -constants.PlayerJumpForce))
    }
}

func (player *Player) handleWalk() {
    if rl.IsKeyDown(rl.KeyD) {
        player.velocityHandler.ModifyVelocity(
            "walk",
            rl.NewVector2(constants.PlayerSpeed, 0))
    
    }

    if rl.IsKeyDown(rl.KeyA) {
        player.velocityHandler.ModifyVelocity(
            "walk",
            rl.NewVector2(-constants.PlayerSpeed, 0))
    }
}

func (player *Player) handleGravity() {
    player.velocityHandler.ModifyVelocity("gravity", rl.NewVector2(0, constants.GameWorldGravity))

    floorHeight := player.gameWorld.GetFloorHeight(player.hitbox.X)

    if player.hitbox.Y + player.hitbox.Height >= float32(floorHeight) && 
        player.velocityHandler.GetVelocityValue("gravity").Value.Y >= 0 {

        player.velocityHandler.SetVelocity("gravity", rl.NewVector2(0, 0))

        player.hitbox.Y = float32(floorHeight) - player.hitbox.Height
    }
}

// Would say "is touching grass", but let's be honest... Gamers don't touch grass..play
func (player *Player) isTouchingGround() bool {
    floorLeeway := float32(20)

    return player.hitbox.Y + player.hitbox.Height >= float32(player.gameWorld.GetFloorHeight(player.hitbox.X)) - floorLeeway
}
