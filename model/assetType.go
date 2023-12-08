package model

import "gorm.io/gorm"

type AssetType struct {
	gorm.Model
	Type string `gorm:"type:ENUM('laptop', 'monitor');default:'laptop'" json:"type"`
}
