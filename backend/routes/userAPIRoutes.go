package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func SetUserAPIPaths(h *handlers.UserHandler, rg *gin.RouterGroup) {
	rg.GET("/user", h.ShowAllUsersHandler)
	rg.GET("/user/:cpf", h.ShowUserByCpfHandler)
	rg.POST("/user", h.CreatUserHandler)
	rg.PUT("/user/status/:cpf", h.UserActivateStatusChangerByCpfHandler)
	rg.PUT("/user/:cpf", h.UserEditByCpf)
}