package models

import (
	"time"
)

type User struct {
	ID        		int       		`json:"id" db:"id"`
	Nome      		string    		`json:"nome" db:"nome"`
	CPF       		string    		`json:"cpf" db:"cpf"`
	SenhaHash		string    		`json:"senha_hash" db:"senha_hash"`
	Equipamentos 	[]Equipamento 	`json:"equipamentos,omitempty"`
	Role      		string    		`json:"role" db:"role"`
	Ativo     		bool      		`json:"ativo" db:"ativo"`
	CreatedAt 		time.Time 		`json:"created_at" db:"created_at"`
}