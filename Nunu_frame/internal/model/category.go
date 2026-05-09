package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID         uint           `gorm:"primaryKey;autoIncrement;comment:分类id" json:"id"`
	Name       string         `gorm:"type:varchar(50);not null;comment:分类名称" json:"name"`
	Remark     string         `gorm:"type:varchar(100);comment:分类描述" json:"remark"`
	Sort       uint           `gorm:"type:int;default:0;comment:排序" json:"sort"`
	Status     uint           `gorm:"type:int;default:0;comment:状态" json:"status"`
	CreateTime time.Time      `gorm:"type:datetime;autoCreateTime;comment:创建时间" json:"create_time"`
	UpdateTime time.Time      `gorm:"type:datetime;autoUpdateTime;comment:更新时间" json:"update_time"`
	DeleteTime gorm.DeletedAt `gorm:"type:datetime;index;comment:删除时间" json:"delete_time"`

	SubCategories []*SubCategory `gorm:"foreignKey:CategoryID;comment:子分类" json:"sub_categories"`
}

func (m *Category) TableName() string {
	return "category"
}
