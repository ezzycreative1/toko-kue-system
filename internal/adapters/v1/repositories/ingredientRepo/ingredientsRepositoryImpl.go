package ingredientRepo

import (
	"context"
	"time"

	"github.com/ezzycreative1/toko-kue-system/internal/core/models/v1/ingredientModel"
	"github.com/ezzycreative1/toko-kue-system/internal/core/ports/v1/ingredientPort"
	"gorm.io/gorm"
)

type IngredientRepository struct {
	DB             *gorm.DB
	KeyTransaction string
	timeout        time.Duration
}

func NewIngredientRepository(db *gorm.DB, keyTransaction string, timeout int) ingredientPort.IIngredientRepository {
	return &IngredientRepository{
		DB:             db,
		KeyTransaction: keyTransaction,
		timeout:        time.Duration(timeout) * time.Second,
	}
}

func (r *IngredientRepository) CreateIngredient(ctx context.Context, input ingredientModel.Ingredients) error {
	trx, ok := ctx.Value(r.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = r.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := trx.WithContext(ctxWT).Create(&input)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
