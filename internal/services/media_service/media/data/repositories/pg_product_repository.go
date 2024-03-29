package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	gormpgsql "github.com/meysamhadeli/shop-golang-microservices/internal/pkg/gorm_pgsql"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/logger"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/utils"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/data/contracts"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/models"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PostgresMediaRepository struct {
	log  logger.ILogger
	cfg  *gormpgsql.GormPostgresConfig
	db   *pgxpool.Pool
	gorm *gorm.DB
}

func NewPostgresMediaRepository(log logger.ILogger, cfg *gormpgsql.GormPostgresConfig, gorm *gorm.DB) contracts.MediaRepository {
	return &PostgresMediaRepository{log: log, cfg: cfg, gorm: gorm}
}

func (p *PostgresMediaRepository) GetAllMedias(ctx context.Context, listQuery *utils.ListQuery) (*utils.ListResult[*models.Media], error) {

	result, err := gormpgsql.Paginate[*models.Media](ctx, listQuery, p.gorm)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *PostgresMediaRepository) SearchMedias(ctx context.Context, searchText string, listQuery *utils.ListQuery) (*utils.ListResult[*models.Media], error) {

	whereQuery := fmt.Sprintf("%s IN (?)", "Name")
	query := p.gorm.Where(whereQuery, searchText)

	result, err := gormpgsql.Paginate[*models.Media](ctx, listQuery, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *PostgresMediaRepository) GetMediaById(ctx context.Context, uuid uuid.UUID) (*models.Media, error) {

	var media models.Media

	if err := p.gorm.First(&media, uuid).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("can't find the media with id %s into the database.", uuid))
	}

	return &media, nil
}

func (p *PostgresMediaRepository) CreateMedia(ctx context.Context, media *models.Media) (*models.Media, error) {

	if err := p.gorm.Create(&media).Error; err != nil {
		return nil, errors.Wrap(err, "error in the inserting media into the database.")
	}

	return media, nil
}

func (p *PostgresMediaRepository) UpdateMedia(ctx context.Context, updateMedia *models.Media) (*models.Media, error) {

	if err := p.gorm.Save(updateMedia).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error in updating media with id %s into the database.", updateMedia.MediaId))
	}

	return updateMedia, nil
}

func (p *PostgresMediaRepository) DeleteMediaByID(ctx context.Context, uuid uuid.UUID) error {

	var media models.Media

	if err := p.gorm.First(&media, uuid).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("can't find the media with id %s into the database.", uuid))
	}

	if err := p.gorm.Delete(&media).Error; err != nil {
		return errors.Wrap(err, "error in the deleting media into the database.")
	}

	return nil
}
