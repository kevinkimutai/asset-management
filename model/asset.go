package model

import "gorm.io/gorm"

type Asset struct {
	gorm.Model
	SerialNumber string    `json:"serial_number"`
	UserID       int       `json:"user_id" gorm:"default:null"`
	User         User      `gorm:"foreignKey:UserID"`
	AssetTypeID  int       `json:"asset_type_id"`
	AssetType    AssetType `gorm:"foreignKey:AssetTypeID"`
	ConditionID  int       `json:"condition_id"`
	Condition    Condition `gorm:"foreignKey:ConditionID"`
	Manufacturer string
	Make         string
	Name         string
	Year         uint
}
