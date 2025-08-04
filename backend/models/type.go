package models

type Type struct {
	ID            	int    	`json:"id" db:"id"`
	Marca         	string 	`json:"marca" db:"marca"`
	Modelo        	string 	`json:"modelo" db:"modelo"`
	RequerMAC     	*bool	`json:"requer_mac" db:"requer_mac"`
	PonMask			string	`json:"pon_mask" db:"pon_mask"`
	Ativo			bool	`json:"ativo" db:"ativo"`
	Minimo			*int	`json:"minimo" db:"minimo"`
	UnidadeMedida 	string 	`json:"unidade_medida" db:"unidade_medida"`
}