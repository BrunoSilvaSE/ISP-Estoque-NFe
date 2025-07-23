package models

import (
	"time"
)

type Client struct {
	ID        	int       	`json:"id" db:"id"`
	Produtos	[]NFItem 	`json:"produtos,omitempty"` // 'omitempty' para não incluir se a lista estiver vazia
	Ativo     	bool      	`json:"ativo" db:"ativo"`
	CreatedAt 	time.Time 	`json:"created_at" db:"created_at"`
}