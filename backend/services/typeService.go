package services

import (
	"backend/models"
	"backend/repositories"

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