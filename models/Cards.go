package models

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	UserId uint
	Name   string
}
