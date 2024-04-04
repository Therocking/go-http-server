package models

import "gorm.io/gorm"

type Product struct {
  gorm.Model

  Name        string  `gorm:"type:varchar(100);not null" json:"name"`
  Description string  `gorm:"type:varchar(100);not null" json:"description"`
  Price       float64 `gorm:"type:varchar(100);not null" json:"price"`
}
