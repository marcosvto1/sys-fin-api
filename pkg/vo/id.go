package vo

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewId() ID {
	return ID(uuid.New())
}

func ParseId(str string) (ID, error) {
	id, err := uuid.Parse(str)
	return ID(id), err
}
