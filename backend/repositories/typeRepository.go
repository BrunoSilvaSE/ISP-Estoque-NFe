package repositories

import (
	"backend/database"
	"backend/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

type TypeRepository struct {
	db *database.DatabaseCliente
}

func NewTypeRepository(db *database.DatabaseCliente) *TypeRepository {
	return &TypeRepository{db: db}
}

// GET
func (r *TypeRepository) FindAllTypes(c *gin.Context) (*[]models.Type, error) {
	var typ []models.Type
	ctx := c.Request.Context()
	query := "SELECT * FROM type ORDER BY marca;"

	rows, err := r.db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("erro na query context: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Type
		if err := rows.Scan(
			&t.ID,
			&t.Marca,
			&t.Modelo,
			&t.RequerMAC,
			&t.Ativo,
			&t.Minimo,
			&t.UnidadeMedida);
			err != nil {return nil, fmt.Errorf("erro na row scan: %w", err)}
		typ = append(typ, t)
	}

	return &typ, nil
}