package models

type UBS struct {
	ID     uint   `gorm:"primary_key"`
	IBGECode string `json:"codigo_ibge"`
	CNES  string `json:"uf"`
	UFCode string `json:"codigo_uf"`
	Logradouro string `json:"logradouro"`
	Bairro string `json:"bairro"`
  }
  