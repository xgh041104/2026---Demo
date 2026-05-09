package model

import (
	"time"

	"gorm.io/gorm"
)

type Label struct {
	ID         uint           `gorm:"primaryKey;autoIncrement;comment:标签id" json:"id"`
	Name       string         `gorm:"type:varchar(50);not null;comment:标签名称" json:"name"`
	Remark     string         `gorm:"type:varchar(100);comment:标签描述" json:"remark"`
	CreateTime time.Time      `gorm:"type:datetime;comment:创建时间;autoCreateTime" json:"create_time"`
	UpdateTime time.Time      `gorm:"type:datetime;comment:更新时间;autoUpdateTime" json:"update_time"`
	DeleteTime gorm.DeletedAt `gorm:"type:datetime;index;comment:删除时间" json:"delete_time"`
}

func (m *Label) TableName() string {
	return "label"
}
