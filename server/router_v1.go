package server

import (
	"github.com/labstack/echo/v4"
)

func routerGroupV1(handler Handler, e *echo.Echo) {
	e.GET("/favicon.ico", func(c echo.Context) error { return nil })

	api := e.Group("/api") // /api

	v1 := api.Group("/v1", func(next echo.HandlerFunc) echo.HandlerFunc { // middleware for /api/v1
		return func(c echo.Context) error {
			c.Set("Version", "v1")
			return next(c)
		}
	})

	v1.GET("/health-check", handler.healtHandler.HealthCheck)

	productGroup := v2.Group("/product")
	productGroup.POST("", handler.productHandler.CreateProductHandler)
	productGroup.PUT("/:id", handler.productHandler.UpdateProductHandler)
	productGroup.DELETE("/:id", handler.productHandler.DeleteProductHandler)
	productGroup.GET("/:id/detail", handler.productHandler.DetailProductHandler)
	productGroup.GET("/list", handler.productHandler.ListProductHandler)
}
