package models

import (
// "github.com/jinzhu/gorm"
)

type User struct {
	ID     int
	Age    int
	Name   string `gorm:"size:255"` // string 的默认长度为255, 使用 tag 可自定义。
	Height int
}
