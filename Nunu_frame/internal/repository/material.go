package repository

import (
	v1 "Nunu_frame/api/v1"
	"Nunu_frame/internal/model"
	"context"
	"strings"
	"time"
)

type MaterialRepository interface {
	GetAuditingMaterialList(c context.Context) ([]*model.Material, error)
	SaveMaterial(c context.Context, req *model.Material) error
	GetMaterialList(c context.Context) ([]*model.Material, error)
	GetLabelNameStr(c context.Context, LabelIdStr string) (string, error)
	HomeMaterialData(c context.Context, status int) (int64, error)
	MaterialUploadTrend(c context.Context, status int) ([]int64, error)
	UpAuditMaterial(c context.Context, req *v1.UpAuditMaterialReq) error
	GetPcCountData(c context.Context, req *v1.PcData) (int64, error)
	GetPcUrlDataList(c context.Context, req *v1.PcData) ([]*model.Material, error)
	DeletePcMaterial(c context.Context, id uint) error
}

func NewMaterialRepository(
	repository *Repository,
) MaterialRepository {
	return &materialRepository{
		Repository: repository,
	}
}

type materialRepository struct {
	*Repository
}

func (r *materialRepository) GetAuditingMaterialList(c context.Context) ([]*model.Material, error) {
	var materials []*model.Material
	err := r.DB(c).Omit("LabelName").
		Where("status = ?", 0).
		Joins("LEFT JOIN `user` ON `material`.creator_id = `user`.id").
		Select("`material`.*, `user`.`name` AS `creator_name`").
		Order("`material`.created_at DESC").
		Find(&materials).Error
	if err != nil {
		return nil, err
	}
	return materials, nil
}

func (r *materialRepository) SaveMaterial(c context.Context, req *model.Material) error {
	err := r.DB(c).Omit("CreatorName", "LabelName").Create(&model.Material{
		Name:       req.Name,
		CreatorId:  req.CreatorId,
		ImageUrl:   req.ImageUrl,
		LabelId:    req.LabelId,
		CategoryId: req.CategoryId,
		Status:     req.Status,
		Remark:     req.Remark,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *materialRepository) GetMaterialList(c context.Context) ([]*model.Material, error) {
	var materials []*model.Material
	err := r.DB(c).Omit("CreatorName", "LabelName").Order("`material`.created_at DESC").Find(&materials).Error
	if err != nil {
		return nil, err
	}
	return materials, nil
}

// GetLabelNameStr
func (r *materialRepository) GetLabelNameStr(c context.Context, LabelIdStr string) (string, error) {
	var LabelStr strings.Builder
	if LabelIdStr == "" {
		return "", nil
	}
	strSlice := strings.Split(LabelIdStr, ",")
	for idx, LabelId := range strSlice {
		var label *model.Label
		err := r.DB(c).Where("id = ?", LabelId).First(&label).Error
		if err != nil {
			return "", err
		}
		if idx > 0 {
			LabelStr.WriteString(",")
		}
		LabelStr.WriteString(label.Name)
	}
	return LabelStr.String(), nil
}

func (r *materialRepository) HomeMaterialData(c context.Context, status int) (int64, error) {
	var total int64
	switch status {
	case 0:
		err := r.DB(c).Table("material").Count(&total).Error
		if err != nil {
			return 0, err
		}
	case 1:
		err := r.DB(c).Table("material").Where("status = ? ", -1).Count(&total).Error
		if err != nil {
			return 0, err
		}
	case 2:
		today := time.Now().Format("2006-01-02")
		err := r.DB(c).Table("material").
			Where("DATE(created_at) = ?", today).
			Count(&total).Error
		if err != nil {
			return 0, err
		}
	case 3:
		err := r.DB(c).Table("material").Where("status = ? ", 1).Count(&total).Error
		if err != nil {
			return 0, err
		}
	}
	return total, nil
}

func (r *materialRepository) MaterialUploadTrend(c context.Context, status int) ([]int64, error) {
	type TempRow struct {
		Day string `gorm:"column:day"`
		Num int64  `gorm:"column:num"`
	}
	loc, _ := time.LoadLocation("Asia/Shanghai") // 强制东八区
	now := time.Now().In(loc)
	// 近7天：从6天前0点到今天
	startDay := now.AddDate(0, 0, -6)
	startDay = time.Date(startDay.Year(), startDay.Month(), startDay.Day(), 0, 0, 0, 0, loc)

	var sqlCase string
	if status == 1 {
		sqlCase = `SUM(CASE 
			WHEN image_url IS NOT NULL 
			AND LOWER(SUBSTRING_INDEX(image_url, '.', -1)) IN ('mp4','avi','mov','flv','wmv') 
			THEN 1 ELSE 0 END) AS num`
	} else {
		sqlCase = `SUM(CASE 
			WHEN image_url IS NOT NULL 
			AND LOWER(SUBSTRING_INDEX(image_url, '.', -1)) IN ('jpg','jpeg','png','gif','webp') 
			THEN 1 ELSE 0 END) AS num`
	}

	var rows []TempRow
	err := r.DB(c).Table("material").
		Where("created_at >= ?", startDay).
		Select("DATE(created_at) AS day, " + sqlCase).
		Group("day").
		Order("day ASC").
		Find(&rows).Error

	if err != nil {
		return nil, err
	}

	dayMap := make(map[string]int64)
	for _, row := range rows {
		// 解析数据库返回的带时区日期，再转成纯日期字符串
		dayTime, _ := time.ParseInLocation("2006-01-02T15:04:05-07:00", row.Day, loc)
		pureDay := dayTime.Format("2006-01-02") // 转成2026-05-06这种格式
		dayMap[pureDay] = row.Num

	}

	result := make([]int64, 7)
	for i := 0; i < 7; i++ {
		day := startDay.AddDate(0, 0, i)
		dateStr := day.Format("2006-01-02")
		result[i] = dayMap[dateStr]
	}
	return result, nil

}

func (r *materialRepository) UpAuditMaterial(c context.Context, req *v1.UpAuditMaterialReq) error {
	Id := req.Id
	Status := req.Status
	remark := req.Remark
	if remark != "" {
		err := r.DB(c).Model(&model.Material{}).
			Where("id = ?", Id).
			Omit("CreatorName", "LabelName").
			Update("status", Status).
			Update("remark", remark).
			Error
		if err != nil {
			return err
		}
	} else {
		err := r.DB(c).Model(&model.Material{}).
			Where("id = ?", Id).
			Omit("CreatorName", "LabelName").
			Update("status", Status).
			Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *materialRepository) GetPcCountData(c context.Context, req *v1.PcData) (int64, error) {
	Id := req.Id
	status := req.Status
	var Count int64
	err := r.DB(c).Table("material").Where("creator_id = ? AND status = ?", Id, status).Count(&Count).Error
	if err != nil {
		return 0, err
	}
	return Count, nil
}

func (r *materialRepository) GetPcUrlDataList(c context.Context, req *v1.PcData) ([]*model.Material, error) {
	Id := req.Id
	status := req.Status
	var materials []*model.Material
	err := r.DB(c).Omit("CreatorName", "LabelName").
		Where("creator_id = ? AND status = ?", Id, status).
		Order("created_at DESC").
		Find(&materials).
		Error
	if err != nil {
		return nil, err
	}
	return materials, nil
}

func (r *materialRepository) DeletePcMaterial(c context.Context, id uint) error {
	err := r.DB(c).Omit("CreatorName", "LabelName").
		Where("id = ?", id).
		Delete(&model.Material{}).
		Error
	if err != nil {
		return err
	}
	return nil
}
