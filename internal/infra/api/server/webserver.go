package server

import (
	"github.com/Genarodaniel/order-system/config/env"
	"github.com/Genarodaniel/order-system/internal/infra/api"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {

	gin.SetMode(env.Config.GinMode)
	r := gin.New()

	r.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/"))
	r.Use(gin.CustomRecovery(PanicFilter))

	api.Router(r)

	return r
}
