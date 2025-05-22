package server

import (
	"github.com/ezzycreative1/toko-kue-system/deps"
	handler "github.com/ezzycreative1/toko-kue-system/internal/adapters/v1/handlers/http"
	handlerV2 "github.com/ezzycreative1/toko-kue-system/internal/adapters/v2/handlers/http"
	"github.com/ezzycreative1/toko-kue-system/pkg/bvalidator"
)

type Handler struct {
	healtHandler handler.CheckHandler
	brandHandler handler.BrandHandler
}

func SetupHandler(dep deps.Dependency) Handler {
	//init validator
	validator := bvalidator.New()

	return Handler{
		healtHandler:      handler.NewHealthCheckHandler(dep.Logger),
		clientTypeHandler: handlerV2.NewClientTypeHandler(dep.ClientTypeService, dep.Logger, validator),
	}
}
