package core

import "github.com/google/uuid"

type Identifier struct {
	id uuid.UUID
}

func (i Identifier) String() string {
	return i.id.String()
}

func NewUuidV4() Identifier {
	return Identifier{
		id: uuid.New(),
	}
}

func NewUuidV7() (Identifier, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return Identifier{}, err
	}
	return Identifier{
		id,
	}, nil
}
