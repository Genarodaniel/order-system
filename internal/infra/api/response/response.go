package response

import (
	"fmt"
	"net/http"
	"strings"
	"unicode"

	"github.com/Genarodaniel/order-system/config/env"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

// Success: return status code ok
func Success(c *gin.Context, content interface{}) {
	c.JSON(http.StatusOK, content)
}

// BadRequest: return status code bad request with message error
func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"message": message})
}

// InternalServerError: return status code internal server error with message error
func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": message})
}

func HasError(c *gin.Context, statusCode int, err error) bool {
	if err != nil {
		log.Error(err)

		isDebug := strings.ToLower(env.Config.LogLevel) == "debug"
		if isDebug && c.Request.Method == http.MethodPost {
			if f := c.Request.Form; f != nil {
				log.Error(fmt.Sprintf("Form %s", f))
			}
		}

		if statusCode == http.StatusInternalServerError {
			message := "Ocorreu um erro interno"
			if isDebug {
				message = err.Error()
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, ResponseError{
				Error:   true,
				Message: UpperFirstLetter(message),
			})
			return true
		}

		c.AbortWithStatusJSON(statusCode, ResponseError{
			Error:   true,
			Message: UpperFirstLetter(err.Error()),
		})
		return true
	}

	return false
}

func UpperFirstLetter(message string) string {
	messageUpper := []rune(message)
	if len(messageUpper) > 0 {
		messageUpper[0] = unicode.ToUpper(messageUpper[0])
	}

	return string(messageUpper)
}
