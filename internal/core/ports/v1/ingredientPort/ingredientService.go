package ingredientPort

import (
	"context"

	"github.com/ezzycreative1/toko-kue-system/internal/core/models/v1/ingredientModel"
)

type IIngredientService interface {
	CreateIngredient(ctx context.Context, input ingredientModel.CreateIngredienyRequest) (int, error)
}
