package dtos

import uuid "github.com/satori/go.uuid"

type CreateMediaResponseDto struct {
	MediaId uuid.UUID `json:"productId"`
}
