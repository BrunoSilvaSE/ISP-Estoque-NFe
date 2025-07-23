package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Equipamento struct {
	UUID          	uuid.UUID      `json:"uuid" db:"uuid"`
	MacID         	sql.NullString `json:"mac_id,omitempty" db:"mac_id"`
	PONSerial     	sql.NullString `json:"pon_serial,omitempty" db:"pon_serial"`
	IDType        	int            `json:"id_type" db:"id_type"`
	Quantidade    	int            `json:"quantidade" db:"quantidade"`
	CustodianteType string         `json:"custodiante_type" db:"custodiante_type"`
	CustodianteID 	sql.NullString `json:"custodiante_id,omitempty" db:"custodiante_id"`
	NFChaveAcesso 	sql.NullString `json:"nf_chave_acesso,omitempty" db:"nf_chave_acesso"`
	Ativo         	bool           `json:"ativo" db:"ativo"`
	CreatedAt     	time.Time      `json:"created_at" db:"created_at"`
}