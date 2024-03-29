package creating_product

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/logger"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/rabbitmq"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/config"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/consumers"
	creatingproductcommandsv1 "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/creating_product/v1/commands"
	creatingproductdtosv1 "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/creating_product/v1/dtos"
	creatingproducteventsv1 "github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/creating_product/v1/events"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/shared/delivery"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/shared/test_fixture"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
)

type createMediaIntegrationTests struct {
	*test_fixture.IntegrationTestFixture
}

var consumer rabbitmq.IConsumer[*delivery.MediaDeliveryBase]

func TestRunner(t *testing.T) {

	var integrationTestFixture = test_fixture.NewIntegrationTestFixture(t, fx.Options(
		fx.Invoke(func(ctx context.Context, jaegerTracer trace.Tracer, log logger.ILogger, connRabbitmq *amqp.Connection, cfg *config.Config) {
			consumer = rabbitmq.NewConsumer(ctx, cfg.Rabbitmq, connRabbitmq, log, jaegerTracer, consumers.HandleConsumeCreateMedia)
			err := consumer.ConsumeMessage(creatingproducteventsv1.MediaCreated{}, nil)
			if err != nil {
				assert.Error(t, err)
			}
		})))

	//https://pkg.go.dev/testing@master#hdr-Subtests_and_Sub_benchmarks
	t.Run("A=create-media-integration-tests", func(t *testing.T) {

		testFixture := &createMediaIntegrationTests{integrationTestFixture}
		testFixture.Test_Should_Create_New_Media_To_DB()
	})
}

func (c *createMediaIntegrationTests) Test_Should_Create_New_Media_To_DB() {

	command := creatingproductcommandsv1.NewCreateMedia(gofakeit.Name(), gofakeit.AdjectiveDescriptive(), gofakeit.Price(150, 6000), 1, 1)
	result, err := mediatr.Send[*creatingproductcommandsv1.CreateMedia, *creatingproductdtosv1.CreateMediaResponseDto](c.Ctx, command)

	assert.NoError(c.T, err)
	assert.NotNil(c.T, result)
	assert.Equal(c.T, command.MediaID, result.MediaId)

	isPublished := c.RabbitmqPublisher.IsPublished(creatingproducteventsv1.MediaCreated{})
	assert.Equal(c.T, true, isPublished)

	isConsumed := consumer.IsConsumed(creatingproducteventsv1.MediaCreated{})
	assert.Equal(c.T, true, isConsumed)

	createdMedia, err := c.IntegrationTestFixture.MediaRepository.GetMediaById(c.Ctx, result.MediaId)
	assert.NoError(c.T, err)
	assert.NotNil(c.T, createdMedia)
}
