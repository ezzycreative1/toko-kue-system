package ingredientModel

type CreateIngredienyRequest struct {
	Name      string  `json:"name,omitempty" validate:"required"`
	Unit      string  `json:"unit,omitempty" validate:"required"`
	PriceUnit float64 `json:"price_per_unit,omitmepty" validate:"required"`
}
