package ingredientPort

import (
	"context"

	"github.com/ezzycreative1/toko-kue-system/internal/core/models/v1/ingredientModel"
)

type IIngredientRepository interface {
	CreateIngredient(ctx context.Context, input ingredientModel.Ingredients) error
}
