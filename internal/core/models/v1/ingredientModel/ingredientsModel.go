package ingredientModel

import (
	"github.com/ezzycreative1/toko-kue-system/internal/core/models/helperModel"
)

type Ingredients struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Unit      string  `json:"unit"`
	PriceUnit float64 `json:"price_per_unit"`
	helperModel.DateAuditModel
	helperModel.UserAuditModel
}

func (a *Ingredients) TableName() string {
	return "public.tmst_ingredients"
}
