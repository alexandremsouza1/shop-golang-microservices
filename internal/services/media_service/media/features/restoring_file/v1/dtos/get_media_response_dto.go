package dtos

import (
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/utils"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/dtos"
)

type GetMediasResponseDto struct {
	Medias *utils.ListResult[*dtos.MediaDto]
}
