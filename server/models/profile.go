package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Name   string `json:"name"`
	Bio    string `json:"bio"`
	Addres string `json:"addres"`
	UserID uint   `gorm:"uniqueIndex"`
}
