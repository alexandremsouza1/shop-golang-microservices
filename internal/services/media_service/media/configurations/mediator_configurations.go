package configurations

import (
	"context"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/grpc"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/logger"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/rabbitmq"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/data/contracts"
	creatingproductv1commands "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/creating_product/v1/commands"
	creatingproductv1dtos "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/creating_product/v1/dtos"
	deletingproductv1commands "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/deleting_product/v1/commands"
	gettingproductbyidv1dtos "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/getting_product_by_id/v1/dtos"
	gettingproductbyidv1queries "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/getting_product_by_id/v1/queries"
	gettingproductsv1dtos "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/getting_products/v1/dtos"
	gettingproductsv1queries "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/getting_products/v1/queries"
	searchingproductv1dtos "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/searching_product/v1/dtos"
	searchingproductv1queries "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/searching_product/v1/queries"
	updatingproductv1commands "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/updating_product/v1/commands"
	updatingproductv1dtos "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/updating_product/v1/dtos"
)

func ConfigMediasMediator(log logger.ILogger, rabbitmqPublisher rabbitmq.IPublisher,
	productRepository contracts.MediaRepository, ctx context.Context, grpcClient grpc.GrpcClient) error {

	//https://stackoverflow.com/questions/72034479/how-to-implement-generic-interfaces
	err := mediatr.RegisterRequestHandler[*creatingproductv1commands.CreateMedia, *creatingproductv1dtos.CreateMediaResponseDto](creatingproductv1commands.NewCreateMediaHandler(log, rabbitmqPublisher, productRepository, ctx, grpcClient))
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler[*gettingproductsv1queries.GetMedias, *gettingproductsv1dtos.GetMediasResponseDto](gettingproductsv1queries.NewGetMediasHandler(log, rabbitmqPublisher, productRepository, ctx, grpcClient))
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler[*searchingproductv1queries.SearchMedias, *searchingproductv1dtos.SearchMediasResponseDto](searchingproductv1queries.NewSearchMediasHandler(log, rabbitmqPublisher, productRepository, ctx, grpcClient))
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler[*updatingproductv1commands.UpdateMedia, *updatingproductv1dtos.UpdateMediaResponseDto](updatingproductv1commands.NewUpdateMediaHandler(log, rabbitmqPublisher, productRepository, ctx, grpcClient))
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler[*deletingproductv1commands.DeleteMedia, *mediatr.Unit](deletingproductv1commands.NewDeleteMediaHandler(log, rabbitmqPublisher, productRepository, ctx, grpcClient))
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler[*gettingproductbyidv1queries.GetMediaById, *gettingproductbyidv1dtos.GetMediaByIdResponseDto](gettingproductbyidv1queries.NewGetMediaByIdHandler(log, rabbitmqPublisher, productRepository, ctx, grpcClient))
	if err != nil {
		return err
	}

	return nil
}
