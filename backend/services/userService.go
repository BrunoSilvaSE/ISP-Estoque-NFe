package services

import (
	"backend/models"
	"backend/repositories"
	"backend/utils"
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService{
	return &UserService{UserRepository: userRepository}
}

// GET
func (s *UserService) SelectAllUsers(c *gin.Context) (*[]models.User, error) {
	return s.UserRepository.FindAllUsers(c)
}

func (s *UserService) SelectUserByCPF(c *gin.Context, cpf string) (*models.User, error) {

	if !utils.IsValidCPF(cpf){
		return nil, errors.New("CPF Inválido")
	}

	return s.UserRepository.FindUserByCPF(c, cpf)
}

// POST
func (s *UserService) UserRegister(c *gin.Context, user *models.User) error {
	
	if !utils.IsValidCPF(user.CPF){
		return errors.New("CPF Inválido")
	}
	
	senha, err := bcrypt.GenerateFromPassword([]byte(user.SenhaHash), bcrypt.DefaultCost)
	if err != nil {
		log.Println("erro in password hash generator: ", err)
		return err
	}
	user.SenhaHash = string(senha)
	
	err = s.UserRepository.InsertUser(c, user)
	if err != nil {
		return err
	}

	return nil
}

// PUT
func (s *UserService) ChangeUserStatusByCpf(c *gin.Context, cpf string) error {
	var newStatus bool

	if !utils.IsValidCPF(cpf){
		return errors.New("CPF Inválido")
	}

	user, err := s.UserRepository.FindUserByCPF(c, cpf)
	if err != nil {
		return fmt.Errorf("erro ao buscar usuário com o cpf %s: %w", cpf, err)
	}

	newStatus = !user.Ativo

	err = s.UserRepository.UserStatusAlterByID(c, newStatus, user.ID)
	if err != nil {
		return fmt.Errorf("erro ao tentar atualizar o status de ativo do usuário %s\n%w", cpf, err)
	}

	return err
}

func (s *UserService) UserUpdateByCpf(c *gin.Context, cpf string, NewUser *models.User) error {
	var err error

	if !utils.IsValidCPF(cpf){
		return errors.New("CPF Inválido")
	}

	existingUser, err := s.UserRepository.FindUserByCPF(c, cpf)
	if err != nil {
		return fmt.Errorf("erro ao buscar usuário com o %s: %w", cpf, err)
	}

	if NewUser.Nome != "" {
		existingUser.Nome = NewUser.Nome
	}

	if NewUser.CPF != "" {
		if !utils.IsValidCPF(cpf){
			return fmt.Errorf("novo CPF Inválido")
		}

		_, err = s.UserRepository.FindUserByCPF(c, NewUser.CPF)
		if err == nil {
			return fmt.Errorf("novo CPF já esta em uso")
		}

		existingUser.CPF = NewUser.CPF
	}

	if NewUser.SenhaHash != "" {
		senha, err := bcrypt.GenerateFromPassword([]byte(NewUser.SenhaHash), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("erro in password hash generator: %w", err)
		}

		existingUser.SenhaHash = string(senha)
	}

	if NewUser.Role != "" {
		existingUser.Role = NewUser.Role
	}

	err = s.UserRepository.UserModifyByID(c, existingUser)
	if err != nil {
		return fmt.Errorf("erro ao tentar modificar usuário \n%w", err)
	}


	return err
}