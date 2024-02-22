package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func KeysHandler(c *gin.Context) {
	if keyResponse, exists := c.Get("keysHandlerResponse"); exists {
		c.JSON(http.StatusOK, keyResponse)
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
}
