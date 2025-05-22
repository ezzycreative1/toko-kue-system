package ingredientService

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ezzycreative1/toko-kue-system/internal/core/models/helperModel"
	"github.com/ezzycreative1/toko-kue-system/internal/core/models/v1/ingredientModel"
	"github.com/ezzycreative1/toko-kue-system/internal/core/ports/v1/ingredientPort"
	"github.com/ezzycreative1/toko-kue-system/pkg/glog"
	"github.com/ezzycreative1/toko-kue-system/pkg/mid"
)

type IngredientService struct {
	IngredientRepo ingredientPort.IIngredientRepository
	Logger         glog.Logger
}

func NewIngredientService(
	ingredientRepo ingredientPort.IIngredientRepository,
	logger glog.Logger,
) ingredientPort.IIngredientService {
	return &IngredientService{
		IngredientRepo: ingredientRepo,
		Logger:         logger,
	}
}

func (s *IngredientService) CreateIngredient(ctx context.Context, input ingredientModel.CreateIngredienyRequest) (int, error) {
	// request id
	requestID := mid.GetIDx(ctx)
	defer s.Logger.InfoT(requestID, "process service completed...", glog.Any("request: ", input))

	// Simpan brand baru ke database
	result := ingredientModel.Ingredients{
		Name:      input.Name,
		Unit:      input.Unit,
		PriceUnit: input.PriceUnit,
		DateAuditModel: helperModel.DateAuditModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UserAuditModel: helperModel.UserAuditModel{
			CreatedBy: "",
			UpdatedBy: "",
		},
	}

	if err := s.IngredientRepo.CreateIngredient(ctx, result); err != nil {
		s.Logger.ErrorT(requestID, fmt.Sprintf("failed to create ingredient: %+v", result), err)
		return http.StatusInternalServerError, fmt.Errorf("failed to create ingredient: %w", err)
	}

	s.Logger.InfoT(requestID, "ingredient created successfully", glog.Any("ingredient", result))
	return http.StatusOK, nil
}
