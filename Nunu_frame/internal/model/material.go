package model

import "gorm.io/gorm"

type Material struct {
	gorm.Model
	Name       string `gorm:"column:name;varchar(50);" json:"name"`
	CreatorId  int    `gorm:"column:creator_id;" json:"creator_id"`
	ImageUrl   string `gorm:"column:image_url;varchar(256);" json:"image_url"`
	LabelId    string `gorm:"column:label_id;varchar(256);" json:"label_id"`
	CategoryId int    `gorm:"column:category_id;" json:"category_id"`
	Status     int    `gorm:"column:status;" json:"status"`
	Remark     string `gorm:"column:remark;varchar(50)" json:"remark"`

	CreatorName string `gorm:"creator_name" json:"creator_name"`
	LabelName   string `gorm:"label_name" json:"label_name"`
}

func (m *Material) TableName() string {
	return "material"
}
