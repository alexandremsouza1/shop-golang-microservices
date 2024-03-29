package commands

import (
	"context"
	"encoding/json"

	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/grpc"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/logger"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/mapper"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/rabbitmq"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/data/contracts"
	dtosv1 "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/creating_product/v1/dtos"
	eventsv1 "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/creating_product/v1/events"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/models"
)

type CreateMediaHandler struct {
	log               logger.ILogger
	rabbitmqPublisher rabbitmq.IPublisher
	productRepository contracts.MediaRepository
	ctx               context.Context
	grpcClient        grpc.GrpcClient
}

func NewCreateMediaHandler(log logger.ILogger, rabbitmqPublisher rabbitmq.IPublisher,
	productRepository contracts.MediaRepository, ctx context.Context, grpcClient grpc.GrpcClient) *CreateMediaHandler {
	return &CreateMediaHandler{log: log, productRepository: productRepository, ctx: ctx, rabbitmqPublisher: rabbitmqPublisher, grpcClient: grpcClient}
}

func (c *CreateMediaHandler) Handle(ctx context.Context, command *CreateMedia) (*dtosv1.CreateMediaResponseDto, error) {

	media := &models.Media{
		MediaId:     command.MediaID,
		Name:        command.Name,
		Description: command.Description,
		Price:       command.Price,
		InventoryId: command.InventoryId,
		Count:       command.Count,
		CreatedAt:   command.CreatedAt,
	}

	createdMedia, err := c.productRepository.CreateMedia(ctx, media)
	if err != nil {
		return nil, err
	}

	evt, err := mapper.Map[*eventsv1.MediaCreated](createdMedia)
	if err != nil {
		return nil, err
	}

	err = c.rabbitmqPublisher.PublishMessage(evt)
	if err != nil {
		return nil, err
	}

	response := &dtosv1.CreateMediaResponseDto{MediaId: media.MediaId}
	bytes, _ := json.Marshal(response)

	c.log.Info("CreateMediaResponseDto", string(bytes))

	return response, nil
}
