package queries

import (
	"context"

	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/grpc"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/logger"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/rabbitmq"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/utils"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/data/contracts"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/dtos"
	dtosv1 "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/getting_products/v1/dtos"
)

type GetMediasHandler struct {
	log               logger.ILogger
	rabbitmqPublisher rabbitmq.IPublisher
	productRepository contracts.MediaRepository
	ctx               context.Context
	grpcClient        grpc.GrpcClient
}

func NewGetMediasHandler(log logger.ILogger, rabbitmqPublisher rabbitmq.IPublisher,
	productRepository contracts.MediaRepository, ctx context.Context, grpcClient grpc.GrpcClient) *GetMediasHandler {
	return &GetMediasHandler{log: log, productRepository: productRepository, ctx: ctx, rabbitmqPublisher: rabbitmqPublisher, grpcClient: grpcClient}
}

func (c *GetMediasHandler) Handle(ctx context.Context, query *GetMedias) (*dtosv1.GetMediasResponseDto, error) {

	products, err := c.productRepository.GetAllMedias(ctx, query.ListQuery)
	if err != nil {
		return nil, err
	}

	listResultDto, err := utils.ListResultToListResultDto[*dtos.MediaDto](products)

	if err != nil {
		return nil, err
	}
	return &dtosv1.GetMediasResponseDto{Medias: listResultDto}, nil
}
