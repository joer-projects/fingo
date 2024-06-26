package core

import "time"

type Entity struct {
	id        string
	updatedBy *string
	updatedAt *time.Time
	createdBy string
	createdAt time.Time
}

type NewEntityProps struct {
	ID        string
	UpdatedBy *string
	UpdatedAt *time.Time
	CreatedBy string
	CreatedAt time.Time
}

type UpdateEntityProps struct {
	UpdatedBy *string
	UpdatedAt *time.Time
}

func NewEntity(props NewEntityProps) Entity {
	return Entity{
		id:        props.ID,
		createdBy: props.CreatedBy,
		createdAt: props.CreatedAt,
	}
}

func (e Entity) GetID() string {
	return e.id
}

func (e Entity) GetUpdatedBy() *string {
	return e.updatedBy
}

func (e Entity) GetUpdatedAt() *time.Time {
	return e.updatedAt
}

func (e Entity) GetCreatedBy() string {
	return e.createdBy
}

func (e Entity) GetCreatedAt() time.Time {
	return e.createdAt
}

func (e Entity) Update(props UpdateEntityProps) Entity {
	return Entity{
		id:        e.id,
		updatedBy: props.UpdatedBy,
		updatedAt: props.UpdatedAt,
		createdBy: e.createdBy,
		createdAt: e.createdAt,
	}
}
