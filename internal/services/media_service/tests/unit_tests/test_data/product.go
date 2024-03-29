package test_data

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/models"
	uuid "github.com/satori/go.uuid"
)

var Medias = []*models.Media{
	{
		MediaId:     uuid.NewV4(),
		Name:        gofakeit.Name(),
		CreatedAt:   time.Now(),
		Description: gofakeit.AdjectiveDescriptive(),
		Price:       gofakeit.Price(100, 1000),
	},
	{
		MediaId:     uuid.NewV4(),
		Name:        gofakeit.Name(),
		CreatedAt:   time.Now(),
		Description: gofakeit.AdjectiveDescriptive(),
		Price:       gofakeit.Price(100, 1000),
	},
}
