package mappings

import (
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/dtos"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/models"
)

func MediaToMediaResponseDto(media *models.Media) *dtos.MediaDto {
	return &dtos.MediaDto{
		MediaId:     media.MediaId,
		Name:        media.Name,
		Description: media.Description,
		Price:       media.Price,
		CreatedAt:   media.CreatedAt,
		UpdatedAt:   media.UpdatedAt,
	}
}
