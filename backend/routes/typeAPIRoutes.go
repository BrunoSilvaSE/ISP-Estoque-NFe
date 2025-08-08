package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func SetTypesAPIPath(h *handlers.TypeHandler, rg *gin.RouterGroup) {
	rg.GET("/type", h.ShowAllTypesHandler)
	rg.GET("/type/:modelOrId", h.ShowTypeByModelOrId)
	rg.POST("/type", h.CreatTypeHandler)
	rg.PUT("/type/status/:modelOrId", h.TypeActivateStatusChangerByModelOrIdHandler)
	rg.PUT("/type/:modelOrId", h.TypeEditByModelOrId)
} 