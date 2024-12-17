package pkg_entities

import (
	"fmt"

	"github.com/google/uuid"
)

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

func NewIdStruct[T any](id string) (T, error) {
	nid, err := NewUUIDFromString(id)
	if err != nil {
		return *new(T), err
	}

	return T(&nid), err
}

