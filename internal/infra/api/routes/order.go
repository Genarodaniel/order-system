package routes

import (
	"github.com/Genarodaniel/order-system/config/dependency"
	"github.com/Genarodaniel/order-system/internal/infra/api/handler"
	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup) {
	handler := handler.NewApiOrderHandler(dependency.Repository.Order)

	g.POST("/", handler.Create)
}
