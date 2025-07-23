package models

import (
	"database/sql"
	"time"
)

type NF struct {
	ChaveAcesso   	string          `json:"chave_acesso" db:"chave_acesso"`
	Numero        	sql.NullString  `json:"numero,omitempty" db:"numero"`
	DataEmissao   	time.Time       `json:"data_emissao" db:"data_emissao"`
	Fornecedor    	sql.NullString  `json:"fornecedor,omitempty" db:"fornecedor"`
	ValorTotal    	sql.NullFloat64 `json:"valor_total,omitempty" db:"valor_total"`
	IDResponsavel 	int             `json:"id_responsavel" db:"id_responsavel"`
	Produtos		[]NFItem 		`json:"produtos,omitempty"` // 'omitempty' para não incluir se a lista estiver vazia
	CreatedAt     	time.Time       `json:"created_at" db:"created_at"`
}