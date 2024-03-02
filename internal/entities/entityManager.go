package entities

type IEntityManager interface {
    GetEntities() []IEntity

    SpawnEntity(entity IEntity)
}

type EntityManager struct {
    entities []IEntity
}

func NewEntityManager() IEntityManager {
    return &EntityManager {}
}

func (entityManager *EntityManager) GetEntities() []IEntity {
    return entityManager.entities
}

func (entityManager *EntityManager) SpawnEntity(entity IEntity) {
    entityManager.entities = append(entityManager.entities, entity)
}
