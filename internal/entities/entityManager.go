package entities

import (
	"errors"
	"fmt"
)

type IEntityManager interface {
    GetEntities() []IEntity
    GetEntityById(id string) (IEntity, error)

    SpawnEntity(id string, entity IEntity)
}

type EntityManager struct {
    entities map[string]IEntity
    duplicateEntityIds map[string]int
}

func NewEntityManager() IEntityManager {
    return &EntityManager {
        entities: make(map[string]IEntity),
        duplicateEntityIds: make(map[string]int),
    }
}

func (entityManager *EntityManager) GetEntities() []IEntity {
    var entities []IEntity

    for _, entity := range entityManager.entities {
        entities = append(entities, entity)
    }
    return entities
}

func (entityManager *EntityManager) SpawnEntity(id string, entity IEntity) {
    if entityManager.entities[id] != nil {
        entityManager.duplicateEntityIds[id]++

        id = fmt.Sprintf("%s-%d", id, entityManager.duplicateEntityIds[id])
    }

    entityManager.entities[id] = entity
}

func (entityManager *EntityManager) GetEntityById(id string) (IEntity, error) {
    if entityManager.entities[id] == nil {
        return nil, errors.New("Entity not found")
    }

    return entityManager.entities[id], nil
}
