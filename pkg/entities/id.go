package pkg_entities

import "github.com/google/uuid"

type UUID = uuid.UUID

func NewUUID() UUID {
	return UUID(uuid.New())
}

func NewUUIDFromString(id string) (*UUID, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return &parsedId, nil
}
