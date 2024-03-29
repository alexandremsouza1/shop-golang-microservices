package mappings

import (
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/mapper"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/dtos"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/creating_product/v1/events"
	events2 "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/updating_product/v1/events"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/models"
)

func ConfigureMappings() error {
	err := mapper.CreateMap[*models.Media, *dtos.MediaDto]()
	if err != nil {
		return err
	}

	err = mapper.CreateMap[*models.Media, *events.MediaCreated]()
	if err != nil {
		return err
	}

	err = mapper.CreateMap[*models.Media, *events2.MediaUpdated]()
	if err != nil {
		return err
	}
	return nil
}
