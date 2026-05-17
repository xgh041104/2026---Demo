package model

import (
	"time"

	"gorm.io/gorm"
)

type Material struct {
	ID         uint           `gorm:"primaryKey;autoIncrement;comment:分类id" json:"id"`
	Name       string         `gorm:"column:name;varchar(50);" json:"name"`
	CreatorId  int            `gorm:"column:creator_id;" json:"creator_id"`
	ImageUrl   string         `gorm:"column:image_url;varchar(256);" json:"image_url"`
	LabelId    string         `gorm:"column:label_id;varchar(256);" json:"label_id"`
	CategoryId int            `gorm:"column:category_id;" json:"category_id"`
	Status     int            `gorm:"column:status;" json:"status"`
	Remark     string         `gorm:"column:remark;varchar(50)" json:"remark"`
	CreateTime time.Time      `gorm:"type:datetime;autoCreateTime;comment:创建时间" json:"create_time"`
	UpdateTime time.Time      `gorm:"type:datetime;autoUpdateTime;comment:更新时间" json:"update_time"`
	DeleteTime gorm.DeletedAt `gorm:"type:datetime;index;comment:删除时间" json:"delete_time"`

	CreatorName string `gorm:"creator_name" json:"creator_name"`
	LabelName   string `gorm:"label_name" json:"label_name"`
}

func (m *Material) TableName() string {
	return "material"
}
