package services

import (
	"backend/models"
	"backend/repositories"
	"fmt"

	"github.com/gin-gonic/gin"
)

type TypeService struct {
	TypeRepository *repositories.TypeRepository
}

func NewTypeService(typeRepository *repositories.TypeRepository) *TypeService {
	return &TypeService{TypeRepository: typeRepository}
}

// GET
func (s *TypeService) SelectAllTypes(c *gin.Context) (*[]models.Type, error) {
	return s.TypeRepository.FindAllTypes(c)
}

//	POST
func (s *TypeService) NewTypeRegister(c *gin.Context, typ *models.Type) error {

	if typ.Marca == "" || typ.Modelo == "" || typ.RequerMAC == nil || typ.Minimo == nil {
		return fmt.Errorf("credenciais incompletas")
	}

	err := s.TypeRepository.InsertNewType(c, typ)
	if err != nil {
		return err
	}

	return nil
}