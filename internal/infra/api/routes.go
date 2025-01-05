package api

import (
	"github.com/Genarodaniel/order-system/config/dependency"
	"github.com/Genarodaniel/order-system/internal/infra/api/handler"
	"github.com/gin-gonic/gin"
)

// @Version 1.0
func Router(e *gin.Engine) {
	handler := handler.NewApiOrderHandler(dependency.Repository.Order)

	// healthCheck.Router(e.Group("/"))

	v1 := e.Group("/api/v1")
	order := v1.Group("/order")

	order.POST("/", handler.Create)

}
