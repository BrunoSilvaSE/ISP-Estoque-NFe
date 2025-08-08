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

func (r *TypeRepository) FindTypeByID(c *gin.Context, id int) (*models.Type, error) {
	var typ models.Type
	ctx := c.Request.Context()
	query := `SELECT * FROM "type" WHERE id = $1;`

	row := r.db.DB.QueryRowContext(ctx, query, id)

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
			return nil, fmt.Errorf("tipo de equipamento com ID %d não encontrado", id)
		}
		return nil, fmt.Errorf("falha ao buscar typo de equipamento por ID: %w", err)
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

//	PUT
func (r *TypeRepository) TypeStatusAlterByID(c *gin.Context, newStatus bool, id int) error {
	query := `UPDATE "type" SET ativo = $1 WHERE id = $2`
	ctx := c.Request.Context()

	res, err := r.db.DB.ExecContext(ctx, query, newStatus, id)
	if err != nil {
		return fmt.Errorf("erro on ExecContext: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("falha ao verificar linhas afetadas após atualização de status: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("tipo de equipamento com ID %d não encontrado para atualização de status", id)
	}

	return err
}

func (r *TypeRepository) TypeModifyByID(c *gin.Context, newType *models.Type) error {
		query := `
        UPDATE "type"
        SET marca = $1, modelo = $2, requer_mac = $3, pon_mask = $4, minimo = $5, unidade_medida = $6
        WHERE id = $7`
	ctx := c.Request.Context()

	res, err := r.db.DB.ExecContext(
		ctx,
		query,
		newType.Marca,
		newType.Modelo,
		newType.RequerMAC,
		newType.PonMask,
		newType.Minimo,
		newType.UnidadeMedida,
		newType.ID,
	)
	if err != nil {
        return fmt.Errorf("falha ao atualizar tipo de equipamento %d no DB: %w", newType.ID, err)
    }

    rowsAffected, err := res.RowsAffected()
    if err != nil {
        return fmt.Errorf("falha ao verificar linhas afetadas após atualização do tipo de equipamento: %w", err)
    }
    if rowsAffected == 0 {
        return fmt.Errorf("tipo de equipamento com ID %d não encontrado para atualização", newType.ID)
    }

    return err
}