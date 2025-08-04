package handlers

import (
	"backend/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TypeHandler struct {
	typeService *services.TypeService
}

func NewTypeHandler(typeService *services.TypeService) *TypeHandler {
	return &TypeHandler{typeService: typeService}
}

// GET
func (h *TypeHandler) ShowAllTypesHandler(c *gin.Context) {

	typ, err := h.typeService.SelectAllTypes(c)
	if err != nil {
		err = fmt.Errorf("erro ao buscar tipo de equipamento\n%w", err)
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	c.JSON(http.StatusOK, typ)
}