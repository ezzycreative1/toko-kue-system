package http

import (
	"fmt"
	"net/http"

	"github.com/ezzycreative1/toko-kue-system/internal/core/models/v1/ingredientModel"
	"github.com/ezzycreative1/toko-kue-system/internal/core/ports/v1/ingredientPort"
	"github.com/ezzycreative1/toko-kue-system/pkg/bvalidator"
	"github.com/ezzycreative1/toko-kue-system/pkg/glog"
	"github.com/ezzycreative1/toko-kue-system/pkg/mid"
	"github.com/ezzycreative1/toko-kue-system/pkg/web"
	"github.com/labstack/echo/v4"
)

type IngredientHandler struct {
	IngredientService ingredientPort.IIngredientService
	Logger            glog.Logger
	Validator         bvalidator.Validator
}

func NewIngredientHandler(
	ingredientService ingredientPort.IIngredientService,
	logger glog.Logger,
	validator bvalidator.Validator,
) IngredientHandler {
	return IngredientHandler{
		IngredientService: ingredientService,
		Logger:            logger,
		Validator:         validator,
	}
}

// CreateIngredientHandler godoc
//
//	@Summary		API Create Ingredient
//	@Description	Create Ingredient
//	@Tags			INgredient
//	@Accept			json
//	@Produce		json
//	@Param			IngredientRequest		body	ingredientModel.CreateIngredienyRequest	true	"Request Parameters"
//	@Router			/v1/ingredients [POST]
func (h *IngredientHandler) CreateIngredientHandler(ctx echo.Context) error {
	requestID := mid.GetID(ctx)
	userCtx := mid.SetIDx(ctx.Request().Context(), requestID)

	h.Logger.InfoT(requestID, "process start request...")
	var payload ingredientModel.CreateIngredienyRequest
	if err := ctx.Bind(&payload); err != nil {
		h.Logger.WarnT(requestID, fmt.Sprintf("create ingredient payload: %s", err.Error()), glog.Any("payload", payload))
		return web.ResponseError(ctx, http.StatusBadRequest, "Bad Request", err)
	}

	// log payload
	h.Logger.InfoT(requestID, "create ingredient request", glog.Any("payload", payload))

	// Validate slice payload
	mapErr, err := h.Validator.Struct(payload)
	if err != nil {
		h.Logger.WarnT(requestID, fmt.Sprintf("Bad Request: %s", err.Error()))
		return web.ResponseErrValidation(ctx, "bad request", mapErr)
	}

	statusCode, err := h.IngredientService.CreateIngredient(userCtx, payload)
	if err != nil {
		h.Logger.ErrorT(requestID, "error create ingredient", err)
		return web.ResponseError(ctx, statusCode, err.Error(), err)
	}

	return web.ResponseFormatter(ctx, statusCode, "Ingredient data successfully created", nil)
}
