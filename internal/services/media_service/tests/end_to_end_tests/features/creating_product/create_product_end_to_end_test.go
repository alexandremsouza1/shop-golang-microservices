package creating_product

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gavv/httpexpect/v2"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/media/features/creating_product/v1/dtos"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/shared/test_fixture"
	"go.uber.org/fx"
)

type createMediaEndToEndTests struct {
	*test_fixture.IntegrationTestFixture
}

func TestRunner(t *testing.T) {

	var endToEndTestFixture = test_fixture.NewIntegrationTestFixture(t, fx.Options())

	//https://pkg.go.dev/testing@master#hdr-Subtests_and_Sub_benchmarks
	t.Run("A=create-media-end-to-end-tests", func(t *testing.T) {

		testFixture := &createMediaEndToEndTests{endToEndTestFixture}
		testFixture.Test_Should_Return_Ok_Status_When_Create_New_Media_To_DB()
	})
}

func (c *createMediaEndToEndTests) Test_Should_Return_Ok_Status_When_Create_New_Media_To_DB() {

	tsrv := httptest.NewServer(c.Echo)
	defer tsrv.Close()

	e := httpexpect.Default(c.T, tsrv.URL)

	request := &dtos.CreateMediaRequestDto{
		Name:        gofakeit.Name(),
		Description: gofakeit.AdjectiveDescriptive(),
		Price:       gofakeit.Price(150, 6000),
		InventoryId: 1,
		Count:       1,
	}

	e.POST("/api/v1/products").
		WithContext(c.Ctx).
		WithJSON(request).
		Expect().
		Status(http.StatusCreated)
}
