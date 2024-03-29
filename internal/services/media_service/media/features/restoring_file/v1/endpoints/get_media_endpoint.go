package endpoints

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	echomiddleware "github.com/meysamhadeli/shop-golang-microservices/internal/pkg/http/echo/middleware"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/logger"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/utils"
	dtosv1 "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/getting_products/v1/dtos"
	queriesv1 "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/getting_products/v1/queries"
)

func MapRoute(validator *validator.Validate, log logger.ILogger, echo *echo.Echo, ctx context.Context) {
	group := echo.Group("/api/v1/products")
	group.GET("", getAllMedias(validator, log, ctx), echomiddleware.ValidateBearerToken())
}

// GetAllMedias
// @Tags Medias
// @Summary Get all media
// @Description Get all products
// @Accept json
// @Produce json
// @Param GetMediasRequestDto query dtos.GetMediasRequestDto false "GetMediasRequestDto"
// @Success 200 {object} dtos.GetMediasResponseDto
// @Security ApiKeyAuth
// @Router /api/v1/products [get]
func getAllMedias(validator *validator.Validate, log logger.ILogger, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {

		listQuery, err := utils.GetListQueryFromCtx(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		request := &dtosv1.GetMediasRequestDto{ListQuery: listQuery}
		if err := c.Bind(request); err != nil {
			log.Warn("Bind", err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		query := queriesv1.NewGetMedias(request.ListQuery)

		queryResult, err := mediatr.Send[*queriesv1.GetMedias, *dtosv1.GetMediasResponseDto](ctx, query)

		if err != nil {
			log.Warnf("GetMedias", err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, queryResult)
	}
}
