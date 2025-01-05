package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func PanicFilter(c *gin.Context, recovered interface{}) {
	if err, ok := recovered.(string); ok {
		log.Error(err)
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Ocorreu um erro interno"})
}
