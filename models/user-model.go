package models

import (
	"gorm.io/gorm"
)

type User struct {
  gorm.Model

  FullName string `gorm:"type:varchar(100);not null" json:"fullname"`
  LastName string `gorm:"type:varchar(100);not null" json:"lastname"`
  Username string `gorm:"type:varchar(100);not null" json:"username"`
  Email string `gorm:"type:varchar(100);not null;unique_index" json:"email"`
  Password string `gorm:"type:varchar(100);not null" json:"password"`
}
