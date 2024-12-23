package model

import (
	"fmt"

	"github.com/google/uuid"
)

type ItemID uuid.UUID

func (id ItemID) String() string {
	return uuid.UUID(id).String()
}

func ItemIDFromString(s string) (ItemID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return ItemID{}, fmt.Errorf("parse: %w", err)
	}
	return ItemID(id), nil
}
