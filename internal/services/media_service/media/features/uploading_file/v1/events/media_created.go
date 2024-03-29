package events

import uuid "github.com/satori/go.uuid"

type MediaCreated struct {
	MediaId     uuid.UUID
	InventoryId int64
	Count       int32
}
