package components

import (
	"errors"

	rl "github.com/gen2brain/raylib-go/raylib"
    "github.com/Ollie-Ave/Project_Kat/internal/shared"
)

type IVelocityHandler interface {
    CreateVelocityType(id string, value *Velocity)

    GetVelocityValue(id string) *Velocity
    GetCombinedVelocity() rl.Vector2

    ModifyVelocity(id string, value rl.Vector2) (*Velocity, error)
    SetVelocity(id string, value rl.Vector2) (*Velocity, error)

    HandleFriction()
}

type VelocityHandler struct {
    velocityTypes map[string]*Velocity
}

type Velocity struct {
    Value rl.Vector2
    TerminalVelocity rl.Vector2
    Friction float32
}

func NewVelocityHandler() IVelocityHandler {
    return &VelocityHandler {
        velocityTypes: make(map[string]*Velocity),
    }
}

func NewVelocity(terminalVelocity rl.Vector2, friction float32) *Velocity {
    return &Velocity {
        Value: rl.NewVector2(0, 0),
        TerminalVelocity: terminalVelocity,
        Friction: friction,
    }
}

func (handler *VelocityHandler) GetCombinedVelocity() rl.Vector2 {
    var combinedVelocity rl.Vector2

    for _, value := range handler.velocityTypes {
        combinedVelocity.X += value.Value.X
        combinedVelocity.Y += value.Value.Y
    }

    return combinedVelocity
}

func (handler *VelocityHandler) HandleFriction() {
    for _, value := range handler.velocityTypes {
        if value.Value.X > 0 {
            value.Value.X = shared.ClampValue(value.Value.X - value.Friction, -value.TerminalVelocity.X, 0)
        } else if value.Value.X < 0 {
            value.Value.X = shared.ClampValue(value.Value.X + value.Friction, 0, value.TerminalVelocity.X)
        }
    }
}

func (handler *VelocityHandler) CreateVelocityType(id string, value *Velocity) {
    handler.velocityTypes[id] = value
}

func (handler *VelocityHandler) GetVelocityValue(id string) *Velocity {
    return handler.velocityTypes[id]
}

func (handler *VelocityHandler) ModifyVelocity(id string, value rl.Vector2) (*Velocity, error) {
    velocityType := handler.velocityTypes[id]

    if velocityType == nil {
        return nil, errors.New("Velocity type not found")
    }

    newXVelocity := velocityType.Value.X + value.X

    if newXVelocity < velocityType.TerminalVelocity.X && 
        newXVelocity > -velocityType.TerminalVelocity.X {

        handler.velocityTypes[id].Value.X = shared.ClampValue(newXVelocity, -velocityType.TerminalVelocity.X, velocityType.TerminalVelocity.X)
    } else if newXVelocity > velocityType.TerminalVelocity.X {
        handler.velocityTypes[id].Value.X = velocityType.TerminalVelocity.X
    
    } else if newXVelocity < -velocityType.TerminalVelocity.X {
        handler.velocityTypes[id].Value.X = -velocityType.TerminalVelocity.X
    }

    newYVelocity := velocityType.Value.Y + value.Y

    if newYVelocity < velocityType.TerminalVelocity.Y && 
        newYVelocity > -velocityType.TerminalVelocity.Y {

        handler.velocityTypes[id].Value.Y = shared.ClampValue(newYVelocity, -velocityType.TerminalVelocity.Y, velocityType.TerminalVelocity.Y)
    } else if newYVelocity > velocityType.TerminalVelocity.Y {
        handler.velocityTypes[id].Value.Y = velocityType.TerminalVelocity.Y
    
    } else if newYVelocity < -velocityType.TerminalVelocity.Y {
        handler.velocityTypes[id].Value.Y = -velocityType.TerminalVelocity.Y
    }


    return nil, nil
}

func (handler *VelocityHandler) SetVelocity(id string, value rl.Vector2) (*Velocity, error) {
    velocityType := handler.velocityTypes[id]

    if velocityType == nil {
        return nil, errors.New("Velocity type not found")
    }

    velocityType.Value = value

    return velocityType, nil
}


