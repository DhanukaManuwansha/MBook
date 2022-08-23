package usercore

import "github.com/google/uuid"

func GenerateId() string {
	newId := uuid.New()
	id := newId.String()

	return id
}
