package model

import (
	"gorm.io/gorm"
)

type User struct {
	// Id       uint      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	gorm.Model
	Account    string `gorm:"type:varchar(50);not null;unique;comment:登录账号" json:"account"`
	Password   string `gorm:"type:varchar(100);not null;comment:登录密码" json:"password"`
	RealName   string `gorm:"type:varchar(50);not null;comment:真实姓名" json:"real_name"`
	Phone      string `gorm:"type:varchar(20);comment:手机号" json:"phone"`
	Email      string `gorm:"type:varchar(100);comment:邮箱" json:"email"`
	Type       string `gorm:"type:varchar(20);not null;comment:用户角色 admin/root/user" json:"type"`
	IsDisabled int8   `gorm:"default:0;comment:是否禁用 0启用 1禁用" json:"is_disabled"`
}

func (u *User) TableName() string {
	return "user" // 如果你想完全按图片标题来，这里写 user
}
