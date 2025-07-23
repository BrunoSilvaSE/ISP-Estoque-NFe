package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Historico struct {
	UUID            uuid.UUID      `json:"uuid" db:"uuid"`
	OrigemType      sql.NullString `json:"origem_type,omitempty" db:"origem_type"`
	OrigemID        sql.NullString `json:"origem_id,omitempty" db:"origem_id"`
	DestinoType     string         `json:"destino_type" db:"destino_type"`
	DestinoID       sql.NullString `json:"destino_id,omitempty" db:"destino_id"`
	IDEquipamento   uuid.UUID      `json:"id_equipamento" db:"id_equipamento"`
	Quantidade      int            `json:"quantidade" db:"quantidade"`
	RegistroDoChamado sql.NullString `json:"registro_do_chamado,omitempty" db:"registro_do_chamado"`
	Motivo          string         `json:"motivo" db:"motivo"`
	Observacao      sql.NullString `json:"observacao,omitempty" db:"observacao"`
	DataMovimentacao time.Time     `json:"data_movimentacao" db:"data_movimentacao"`
	IDResponsavel   int            `json:"id_responsavel" db:"id_responsavel"`
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
}