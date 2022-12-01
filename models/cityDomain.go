package models

type City struct {
	ID     uint   `gorm:"primary_key"`
	IBGECode uint `json:"codigo_ibge"`
	UF  string `json:"UF"`
	UFCode string `json:"codigo_UF"`
	CityName string `json:"municipio"`
  }
  