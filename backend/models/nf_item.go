package models

import (
	"time"

	"github.com/google/uuid"
)

type NFItem struct {
	ID            	uuid.UUID 	`json:"id" db:"id"`
	NFChaveAcesso 	string    	`json:"nf_chave_acesso" db:"nf_chave_acesso"`
	IDType        	int       	`json:"id_type" db:"id_type"`
	Quantidade    	int       	`json:"quantidade" db:"quantidade"`
	CreatedAt     	time.Time 	`json:"created_at" db:"created_at"`
}