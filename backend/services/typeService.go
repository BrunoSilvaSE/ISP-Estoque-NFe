package services

import (
	"backend/models"
	"backend/repositories"
	"fmt"
	"strings"

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

func (s *TypeService) SelectTypeByModel(c *gin.Context, model string) (*models.Type, error) {
	model = strings.ToUpper(model)
	return s.TypeRepository.FindTypeByModel(c, model)
}

//	POST
func (s *TypeService) NewTypeRegister(c *gin.Context, typ *models.Type) error {

	if typ.Marca == "" || typ.Modelo == "" || typ.RequerMAC == nil || typ.Minimo == nil {
		return fmt.Errorf("credenciais incompletas")
	}

	typ.Marca = strings.ToUpper(typ.Marca)
	typ.Modelo = strings.ToUpper(typ.Modelo)

	err := s.TypeRepository.InsertNewType(c, typ)
	if err != nil {
		return err
	}

	return nil
}