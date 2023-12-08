package model

import "gorm.io/gorm"

type Asset struct {
	gorm.Model
	SerialNumber string
	UserID       int
	User         User `gorm:"foreignKey:UserID"`
	AssetTypeID  int
	AssetType    AssetType `gorm:"foreignKey:AssetTypeID"`
	ConditionID  int
	Condition    Condition `gorm:"foreignKey:ConditionID"`
	Manufacturer string
	Make         string
	Name         string
	Year         uint
}
