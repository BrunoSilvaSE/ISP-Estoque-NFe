package handlers

import (
	"backend/models"
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

func (h *TypeHandler) ShowTypeByModel(c *gin.Context) {
	model := c.Param("model")

	user, err := h.typeService.SelectTypeByModel(c, model)
	if err != nil {
		err = fmt.Errorf("erro ao buscar usuários.\n%w", err)
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	c.JSON(http.StatusOK, user)
}

// POST
func (h *TypeHandler) CreatTypeHandler(c *gin.Context){
	var typ models.Type

	if err := c.ShouldBindJSON(&typ); err != nil {
		err = fmt.Errorf("erro ao fazer bind do JSON: %w", err)
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	if err := h.typeService.NewTypeRegister(c, &typ); err != nil {
		err = fmt.Errorf("erro ao cadastrar novo tipo de equipamento.\n%w", err)
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensagem": "Novo tipo de equipamento criado com sucesso",
	})
}