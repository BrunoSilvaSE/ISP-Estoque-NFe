package handlers

import (
	"backend/models"
	"backend/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GET
func (h *UserHandler) ShowAllUsersHandler(c *gin.Context) {

	user, err := h.userService.SelectAllUsers(c)
	if err != nil {
		err = fmt.Errorf("erro ao buscar usuários: %w", err)
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) ShowUserByCpfHandler(c *gin.Context) {
	cpf := c.Param("cpf")

	user, err := h.userService.SelectUserByCPF(c, cpf)
	if err != nil {
		err = fmt.Errorf("erro ao buscar usuários \n%w", err)
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	c.JSON(http.StatusOK, user)
}

// POST
func (h *UserHandler) CreatUserHandler(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err = fmt.Errorf("erro ao fazer bind do JSON: %w", err)
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	if err := h.userService.UserRegister(c, &user); err != nil {
		err = fmt.Errorf("erro ao cadastrar usuário: %w", err)
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensagem": "Usuário criado com sucesso",
	})
}

// PUT
func (h *UserHandler) UserActivateStatusChangerByCpfHandler(c *gin.Context) {
	cpf := c.Param("cpf")

	err := h.userService.ChangeUserStatusByCpf(c, cpf)
	if err != nil {
		err = fmt.Errorf("erro ao alterar usuarío \n%w", err)
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Status do usuário alterado com sucesso"})
}

func (h *UserHandler) UserEditByCpf(c *gin.Context) {
	cpf := c.Param("cpf")
	var NewUser models.User

	if err := c.ShouldBindJSON(&NewUser); err != nil {
		err = fmt.Errorf("erro ao fazer bind do JSON: %w", err)
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	err := h.userService.UserUpdateByCpf(c, cpf, &NewUser)
	if err != nil {
		err = fmt.Errorf("erro ao alterar usuarío %s\n%w", cpf, err)
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Usuário alterado com sucesso"})
}