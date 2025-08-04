package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func SetTypesAPIPath(h *handlers.TypeHandler, rg *gin.RouterGroup) {
	rg.GET("/type", h.ShowAllTypesHandler)
	rg.POST("/type", h.CreatTypeHandler)
} 