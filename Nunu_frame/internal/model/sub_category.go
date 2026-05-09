package model

import (
	"time"

	"gorm.io/gorm"
)

type SubCategory struct {
	ID         int            `gorm:"primaryKey;autoIncrement;comment:子分类id" json:"id"`
	Name       string         `gorm:"type:varchar(50);not null;comment:子分类名称" json:"name"`
	Sort       int            `gorm:"type:int;default:0;comment:排序" json:"sort"`
	Status     int            `gorm:"type:int;default:0;comment:状态" json:"status"`
	Remark     string         `gorm:"type:varchar(100);comment:子分类描述" json:"remark"`
	CreateTime time.Time      `gorm:"type:datetime;autoCreateTime;comment:创建时间" json:"create_time"`
	UpdateTime time.Time      `gorm:"type:datetime;autoUpdateTime;comment:更新时间" json:"update_time"`
	DeleteTime gorm.DeletedAt `gorm:"type:datetime;index;comment:删除时间" json:"delete_time"`
	CategoryId int            `gorm:"type:int;comment:分类id" json:"category_id"`
}

func (m *SubCategory) TableName() string {
	return "sub_category"
}
