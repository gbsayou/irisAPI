package models

import (
// "github.com/jinzhu/gorm"
)

type Token struct {
	ID     int
	UserID int
	Token  string `gorm:"size:50"`
	Status string `gorm:"size:1"`
}
