package dtos

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type MediaDto struct {
	MediaId     uuid.UUID `json:"productId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	InventoryId int64     `json:"inventoryId"`
	Count       int32     `json:"count"`
}
