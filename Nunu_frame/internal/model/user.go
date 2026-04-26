package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account  string `gorm:"column:account;varchar(50);not null" json:"account"`
	Password string `gorm:"column:password;varchar(255);not null" json:"password"`
	Type     uint   `gorm:"column:type;not null" json:"type"`
}

func (u *User) TableName() string {
	return "users"
}
