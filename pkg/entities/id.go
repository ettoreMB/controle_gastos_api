package pkg_entities

import "github.com/google/uuid"

type UUID = uuid.UUID

func NewUUID() UUID {
	return UUID(uuid.New())
}
