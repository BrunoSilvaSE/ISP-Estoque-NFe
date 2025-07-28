package repositories

import (
	"backend/database"
	"backend/models"
	"database/sql"
	
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type UserRepository struct {
	db *database.DatabaseCliente
}

func NewUserRepository(db *database.DatabaseCliente) *UserRepository {
	return &UserRepository{db: db}
}

// GET
func (r *UserRepository) FindAllUsers(c *gin.Context) (*[]models.User, error) {
	var users []models.User
	ctx := c.Request.Context()
	query := "SELECT * FROM User ORDER BY nome;"

	rows, err := r.db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("erro na query context: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var u models.User
		if err := rows.Scan(
			&u.ID,
			&u.Nome,
			&u.CPF,
			&u.SenhaHash,
			&u.Role);
			err != nil {return nil, fmt.Errorf("erro na row scan: %w", err)}
		users = append(users, u)
	}

	return &users, nil
}

func (r *UserRepository) FindUserByCPF(c *gin.Context, cpf string) (*models.User, error) {
	var user models.User
	var equipamentos []models.Equipamento
	ctx := c.Request.Context()
	query := "SELECT * FROM User WHERE cpf = $1;"

	row, err := r.db.DB.QueryContext(ctx, query, cpf)
	if err != nil {
		return nil, fmt.Errorf("erro na query context: %w", err)
	}
	defer row.Close()

	err = row.Scan(
		&user.ID,
		&user.Nome,
		&user.CPF,
		&user.SenhaHash,
		&user.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuário com CPF %s não encontrado", cpf)
		}
		return nil, fmt.Errorf("falha ao buscar usuário por ID: %w", err)
	}

	query = "SELECT * FROM Equipamento WHERE custodiante_id = $1;"

	rows, err := r.db.DB.QueryContext(ctx, query, user.ID)
	if err != nil {
		return nil, fmt.Errorf("erro na query context: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var e models.Equipamento
		if err := rows.Scan(
			&e.UUID,
			&e.MacID,
			&e.PONSerial,
			&e.IDType,
			&e.Quantidade,
			&e.CustodianteType,
			&e.CustodianteID,
			&e.NFChaveAcesso,
			&e.Ativo,
			&e.CreatedAt);
			err != nil {
				if err != sql.ErrNoRows {
					return nil, fmt.Errorf("erro falha ao buscar usuários: %w", err)
				}
				break
			}
		equipamentos = append(equipamentos, e)
	}

	user.Equipamentos = equipamentos

	return &user, nil
}

// POST
func (r *UserRepository) InsertUser(c *gin.Context, user *models.User) error {
	query := `INSERT INTO User (id, nome, cpf, senha_hash, role) VALUES ($1, $2, $3, $4, $5)`
	ctx := c.Request.Context()

	res, err := r.db.DB.ExecContext(ctx, query, user.ID, user.Nome, user.CPF, user.SenhaHash, user.Role)
	if err != nil {
		log.Println("erro on ExecContext: ", err)
	}

	if res != nil {
		rowsaffected, err := res.RowsAffected();
		if  err != nil {log.Println("Rows Affected: ", rowsaffected)}
		if rowsaffected == 0 {return sql.ErrNoRows}
	}
	
	return err
}

// PUT
func (r *UserRepository) UserStatusAlterByID(c *gin.Context, newStatus bool, id int) error {
	query := `UPDATE "User" SET ativo = $1 WHERE id = $2`
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
		return fmt.Errorf("usuário com ID %d não encontrado para atualização de status", id)
	}

	return err
}

func (r *UserRepository) UserModifyByID(c *gin.Context, newUser *models.User) error {
	query := `
        UPDATE "User"
        SET nome = $1, cpf = $2, senha_hash = $3, role = $4
        WHERE id = $5`
	ctx := c.Request.Context()

	res, err := r.db.DB.ExecContext(
		ctx,
		query,
		newUser.Nome,
		newUser.CPF,
		newUser.SenhaHash,
		newUser.Role,
		newUser.ID,
	)
	if err != nil {
        return fmt.Errorf("falha ao atualizar usuário %d no DB: %w", newUser.ID, err)
    }

    rowsAffected, err := res.RowsAffected()
    if err != nil {
        return fmt.Errorf("falha ao verificar linhas afetadas após atualização de usuário: %w", err)
    }
    if rowsAffected == 0 {
        return fmt.Errorf("usuário com ID %d não encontrado para atualização", newUser.ID)
    }

    return err
}