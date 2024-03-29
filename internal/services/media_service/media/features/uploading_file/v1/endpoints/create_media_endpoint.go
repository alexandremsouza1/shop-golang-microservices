package endpoints

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	echomiddleware "github.com/meysamhadeli/shop-golang-microservices/internal/pkg/http/echo/middleware"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/logger"
	commandsv1 "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/creating_product/v1/commands"
	dtosv1 "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/creating_product/v1/dtos"
	"github.com/pkg/errors"
)

func MapRoute(validator *validator.Validate, log logger.ILogger, echo *echo.Echo, ctx context.Context) {
	group := echo.Group("/api/v1/products")
	group.POST("", createMedia(validator, log, ctx), echomiddleware.ValidateBearerToken())
}

// CreateMedia
// @Tags        Medias
// @Summary     Create media
// @Description Create new media item
// @Accept      json
// @Produce     json
// @Param       CreateMediaRequestDto body     dtos.CreateMediaRequestDto true "Media data"
// @Success     201                     {object} dtos.CreateMediaResponseDto
// @Security ApiKeyAuth
// @Router      /api/v1/products [post]
func createMedia(validator *validator.Validate, log logger.ILogger, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {

		request := &dtosv1.CreateMediaRequestDto{}

		if err := c.Bind(request); err != nil {
			badRequestErr := errors.Wrap(err, "[createMediaEndpoint_handler.Bind] error in the binding request")
			log.Error(badRequestErr)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		command := commandsv1.NewCreateMedia(request.Name, request.Description, request.Price, request.InventoryId, request.Count)

		if err := validator.StructCtx(ctx, command); err != nil {
			validationErr := errors.Wrap(err, "[createMediaEndpoint_handler.StructCtx] command validation failed")
			log.Error(validationErr)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		result, err := mediatr.Send[*commandsv1.CreateMedia, *dtosv1.CreateMediaResponseDto](ctx, command)

		if err != nil {
			log.Errorf("(CreateMedia.Handle) id: {%s}, err: {%v}", command.MediaID, err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		log.Infof("(media created) id: {%s}", command.MediaID)
		return c.JSON(http.StatusCreated, result)
	}
}
