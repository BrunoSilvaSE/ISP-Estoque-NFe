package services

import (
	"backend/models"
	"backend/repositories"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type TypeService struct {
	TypeRepository *repositories.TypeRepository
}

func NewTypeService(typeRepository *repositories.TypeRepository) *TypeService {
	return &TypeService{TypeRepository: typeRepository}
}

//	GET
func (s *TypeService) SelectAllTypes(c *gin.Context) (*[]models.Type, error) {
	return s.TypeRepository.FindAllTypes(c)
}

func (s *TypeService) SelectTypeByModelOrId(c *gin.Context, modelOrId string) (*models.Type, error) {
	id, err := strconv.Atoi(modelOrId)
	if err != nil {
		model := strings.ToUpper(modelOrId)
		return s.TypeRepository.FindTypeByModel(c, model)
	} else {
		return s.TypeRepository.FindTypeByID(c, id)
	}
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

//	PUT
func (s *TypeService) ChangeTypeStatusByModelOrId(c *gin.Context, modelOrId string) error {
	var newStatus bool

	typ, err := s.TypeRepository.FindTypeByModel(c, modelOrId)
	if err != nil {
		return fmt.Errorf("erro ao buscar tipo de equipamento modelo %s.\n%w", modelOrId, err)
	}

	newStatus = !typ.Ativo

	err = s.TypeRepository.TypeStatusAlterByID(c, newStatus, typ.ID)
	if err != nil {
		return fmt.Errorf("erro ao tentar atualizar o status de ativo do modelo %s\n%w", typ.Modelo, err)
	}

	return err
}

func (s *TypeService) TypeUpdateByModelOrId(c *gin.Context,newType *models.Type, modelOrId string) error {

	OldType, err := s.SelectTypeByModelOrId(c, modelOrId)

	if err != nil {
		return fmt.Errorf("erro ao buscar tipo de modelo\n%w", err)
	}

	if newType.Marca == ""{
		newType.Marca = strings.ToUpper(OldType.Marca)
	} else {
		newType.Marca = strings.ToUpper(newType.Marca)
	}

	if newType.Modelo == ""{
		newType.Modelo = strings.ToUpper(OldType.Modelo)
	} else {
		newType.Modelo = strings.ToUpper(newType.Modelo)
	}

	if newType.RequerMAC == nil{
		newType.RequerMAC = OldType.RequerMAC
	}

	if newType.Minimo == nil{
		newType.Minimo = OldType.Minimo
	}

	if newType.UnidadeMedida == ""{
		newType.UnidadeMedida = OldType.UnidadeMedida
	}

	newType.ID = OldType.ID

	err = s.TypeRepository.TypeModifyByID(c, newType)
	if err != nil{
		return fmt.Errorf("erro ao tentar alterar o tipo de equipamento pelo ID\n%w", err)
	}

	return err
}