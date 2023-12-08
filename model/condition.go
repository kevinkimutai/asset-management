package model

import "gorm.io/gorm"

type Condition struct {
	gorm.Model
	Condition string `gorm:"type:ENUM('new', 'pre-owned');default:'new'" json:"condition"`
}
