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

//	GET
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

func (h *TypeHandler) ShowTypeByModelOrId(c *gin.Context) {
	modelOrId := c.Param("modelOrId")

	user, err := h.typeService.SelectTypeByModelOrId(c, modelOrId)
	if err != nil {
		err = fmt.Errorf("erro ao buscar usuários.\n%w", err)
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	c.JSON(http.StatusOK, user)
}

//	POST
func (h *TypeHandler) CreatTypeHandler(c *gin.Context) {
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

//	PUT
func (h *TypeHandler) TypeActivateStatusChangerByModelOrIdHandler(c *gin.Context) {
	modelOrId := c.Param("modelOrId")

	err := h.typeService.ChangeTypeStatusByModelOrId(c, modelOrId)
	if err != nil {
		err = fmt.Errorf("erro ao alterar tipo de equipamento.\n%w", err)
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Status do usuário alterado com sucesso"})
}

func (h *TypeHandler) TypeEditByModelOrId(c *gin.Context) {
	var newType models.Type
	modelOrId := c.Param("modelOrId")

	if err := c.ShouldBindJSON(&newType); err != nil {
		err = fmt.Errorf("erro ao fazer bind do JSON: %w", err)
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	err := h.typeService.TypeUpdateByModelOrId(c, &newType, modelOrId)
	if err != nil{
		err = fmt.Errorf("erro ao tentar editar tipo de inte.\n%w",err)
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Dados do tipo de item atualizado com sucesso"})
}