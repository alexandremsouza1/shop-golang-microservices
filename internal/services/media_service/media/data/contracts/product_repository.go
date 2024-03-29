package contracts

import (
	"context"

	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/utils"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/models"

	uuid "github.com/satori/go.uuid"
)

type MediaRepository interface {
	GetAllMedias(ctx context.Context, listQuery *utils.ListQuery) (*utils.ListResult[*models.Media], error)
	SearchMedias(ctx context.Context, searchText string, listQuery *utils.ListQuery) (*utils.ListResult[*models.Media], error)
	GetMediaById(ctx context.Context, uuid uuid.UUID) (*models.Media, error)
	CreateMedia(ctx context.Context, media *models.Media) (*models.Media, error)
	UpdateMedia(ctx context.Context, media *models.Media) (*models.Media, error)
	DeleteMediaByID(ctx context.Context, uuid uuid.UUID) error
}
