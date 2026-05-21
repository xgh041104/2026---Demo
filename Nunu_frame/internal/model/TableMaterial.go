package model

import (
	"time"

	"gorm.io/gorm"
)

type TableMaterial struct {
    ID          uint           `gorm:"primaryKey;autoIncrement;comment:分类id" json:"id"`
    Name        string         `gorm:"column:name;type:varchar(50);comment:素材主题" json:"name"`
    CreatorId   int            `gorm:"column:creator_id;type:int;comment:上传者id" json:"creator_id"`
    ImageUrl    string         `gorm:"column:image_url;type:varchar(256);comment:文件路径" json:"image_url"`
    LabelId     string         `gorm:"column:label_id;type:varchar(256);comment:标签id" json:"label_id"`
    CategoryId  int            `gorm:"column:category_id;type:int;comment:分类id" json:"category_id"`
    Status      int            `gorm:"column:status;type:int;comment:上传状态" json:"status"`
    Remark      string         `gorm:"column:remark;type:varchar(50);comment:备注" json:"remark"`
    CreateTime  time.Time      `gorm:"type:datetime;autoCreateTime;comment:创建时间" json:"create_time"`
    UpdateTime  time.Time      `gorm:"type:datetime;autoUpdateTime;comment:更新时间" json:"update_time"`
    DeleteTime  gorm.DeletedAt `gorm:"type:datetime;index;comment:删除时间" json:"delete_time"`
}

func (m *TableMaterial) TableName() string {
	return "material"
}
