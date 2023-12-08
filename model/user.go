package model

import "gorm.io/gorm"

type Role string

// const (
// 	superadmin Role = "admin"
// 	admin      Role = "superadmin"
// 	user       Role = "user"
// )

type User struct {
	gorm.Model
	FirstName   string `json:"first_name" gorm:"not null"`
	LastName    string `json:"last_name" gorm:"not null"`
	Email       string `json:"email" gorm:"not null unique"`
	Designation string
	Role        string `gorm:"type:ENUM('superadmin', 'admin', 'user');default:'user'" json:"role"`
	Asset       []Asset
	Password    string
}
