//go:build integration

package handlers_test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Aritiaya50217/GolangByCodeBangkok/handlers"
	"github.com/Aritiaya50217/GolangByCodeBangkok/repositories"
	"github.com/Aritiaya50217/GolangByCodeBangkok/services"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscountIntegrationService(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		amount := 100
		expected := 80

		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{
			ID:              1,
			PurchaseMin:     100,
			DiscountPercent: 20,
		}, nil)

		promoService := services.NewPromotionService(promoRepo)
		promoHandler := handlers.NewPromotionHandler(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		// Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		// Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expected), string(body))
		}
	})
}
