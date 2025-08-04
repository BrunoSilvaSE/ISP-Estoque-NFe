package repositories

import (
	"backend/database"
	"backend/models"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

type TypeRepository struct {
	db *database.DatabaseCliente
}

func NewTypeRepository(db *database.DatabaseCliente) *TypeRepository {
	return &TypeRepository{db: db}
}

//	GET
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
			&t.PonMask,
			&t.Ativo,
			&t.Minimo,
			&t.UnidadeMedida);
			err != nil {return nil, fmt.Errorf("erro na row scan: %w", err)}
		typ = append(typ, t)
	}

	return &typ, nil
}

func (r *TypeRepository) FindTypeByModel(c *gin.Context, model string) (*models.Type, error) {
	var typ models.Type
	ctx := c.Request.Context()
	query := `SELECT * FROM "type" WHERE modelo = $1;`

	row := r.db.DB.QueryRowContext(ctx, query, model)

	err := row.Scan(
		&typ.ID,
		&typ.Marca,
		&typ.Modelo,
		&typ.RequerMAC,
		&typ.PonMask,
		&typ.Ativo,
		&typ.Minimo,
		&typ.UnidadeMedida)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tipo de equipamento com modelo %s não encontrado", model)
		}
		return nil, fmt.Errorf("falha ao buscar typo de equipamento por modelo: %w", err)
	}

	return &typ, nil
}

//	POST
func (r *TypeRepository) InsertNewType(c *gin.Context, typ *models.Type) error {
	query := `INSERT INTO "type" (marca, modelo, requer_mac, pon_mask, minimo, unidade_medida) VALUES ($1, $2, $3, $4, $5, $6)`
	ctx := c.Request.Context()

	res, err := r.db.DB.ExecContext(ctx, query, typ.Marca, typ.Modelo, typ.RequerMAC, typ.PonMask, typ.Minimo, typ.UnidadeMedida)
	if err != nil {
		err = fmt.Errorf("erro on ExecContext: %w", err)
		return err
	}

	if res != nil {
		rowsaffected, err := res.RowsAffected();
		if  err != nil {
			err = fmt.Errorf("rows Affected: %v", rowsaffected)
			return err
		}
		if rowsaffected == 0 {return sql.ErrNoRows}
	}
	
	return err
}