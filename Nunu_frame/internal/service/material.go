package service

import (
	v1 "Nunu_frame/api/v1"
	"path/filepath"
	"strings"

	"Nunu_frame/internal/model"
	"Nunu_frame/internal/repository"
	"context"
)

type MaterialService interface {
	GetAuditingMaterialList(c context.Context, req *v1.PageNumAndPageSizeRequest) (*v1.ReqAuditingMaterialListResponse, error)
	SaveAvatar(ctx context.Context, req *model.Material) error
	GetMaterialList(c context.Context, req *v1.ReqSearchMaterial) (*v1.ReqMaterialListRes, error)
	HomeMaterialData(c context.Context, status int) (int64, error)
	MaterialUploadTrend(c context.Context) (*v1.ReqListListInt, error)
	UpAuditMaterial(c context.Context, req *v1.UpAuditMaterialReq) error
	GetPcCountData(c context.Context, req *v1.PcData) (int64, error)
	GetPcUrlDataList(c context.Context, req *v1.PcData) (*v1.ResPcDataList, error)
	DeletePcMaterial(c context.Context, req *v1.PcData) error
}

func NewMaterialService(
	service *Service,
	materialRepository repository.MaterialRepository,
) MaterialService {
	return &materialService{
		Service:            service,
		materialRepository: materialRepository,
	}
}

type materialService struct {
	*Service
	materialRepository repository.MaterialRepository
}

func (s *materialService) GetAuditingMaterialList(c context.Context, req *v1.PageNumAndPageSizeRequest) (*v1.ReqAuditingMaterialListResponse, error) {
	materials, err := s.materialRepository.GetAuditingMaterialList(c)
	if err != nil {
		return nil, v1.ErrGetMaterialList
	}

	resp := &v1.ReqAuditingMaterialListResponse{
		ReqMaterialList: make([]v1.ReqAuditingMaterial, 0, len(materials)),
	}

	for _, material := range materials {
		resp.ReqMaterialList = append(resp.ReqMaterialList, v1.ReqAuditingMaterial{
			Id:          material.ID,
			Name:        material.Name,
			CreatorName: material.CreatorName,
			CreatedAt:   material.CreateTime.Format("2006-01-02 15:04:05"),
			ImageUrl:    material.ImageUrl,
		})
	}
	return resp, nil
}

// 存储图片到本地
func (s *materialService) SaveAvatar(ctx context.Context, req *model.Material) error {
	err := s.materialRepository.SaveMaterial(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (s *materialService) GetMaterialList(c context.Context, req *v1.ReqSearchMaterial) (*v1.ReqMaterialListRes, error) {
	materials, err := s.materialRepository.GetMaterialList(c, req)
	if err != nil {
		return nil, v1.ErrGetMaterialList
	}

	resp := &v1.ReqMaterialListRes{
		FileStatus: req.FileStatus,
		ReqMaterialResList: make([]v1.ReqMaterialRes, 0, len(materials)),
	}

	var (
		imageExts = map[string]bool{
			".jpg":  true,
			".jpeg": true,
			".png":  true,
			".gif":  true,
			".bmp":  true,
			".webp": true,
			".svg":  true,
		}
		videoExts = map[string]bool{
			".mp4":  true,
			".avi":  true,
			".mov":  true,
			".wmv":  true,
			".flv":  true,
			".mkv":  true,
			".webm": true,
		}
	)

	for _, material := range materials {
		ext := strings.ToLower(filepath.Ext(material.ImageUrl))
		if imageExts[ext] && req.FileStatus == 0 { // 为图片时
			LabelNameStr, err := s.materialRepository.GetLabelNameStr(c, material.LabelId)
			if err != nil {
				return nil, v1.ErrGetMaterialLabelStr
			}
			resp.ReqMaterialResList = append(resp.ReqMaterialResList, v1.ReqMaterialRes{
				Id:         material.ID,
				Name:       material.Name,
				CreatedAt:  material.CreateTime.Format("2006-01-02 15:04:05"),
				CategoryId: material.CategoryId,
				ImageUrl:   material.ImageUrl,
				LabelName:  LabelNameStr,
			})
		} else if videoExts[ext] && req.FileStatus == 1 { // 为视频时
			LabelNameStr, err := s.materialRepository.GetLabelNameStr(c, material.LabelId)
			if err != nil {
				return nil, v1.ErrGetMaterialLabelStr
			}
			resp.ReqMaterialResList = append(resp.ReqMaterialResList, v1.ReqMaterialRes{
				Id:         material.ID,
				Name:       material.Name,
				CreatedAt:  material.CreateTime.Format("2006-01-02 15:04:05"),
				CategoryId: material.CategoryId,
				ImageUrl:   material.ImageUrl,
				LabelName:  LabelNameStr,
			})
		}
	}
	return resp, nil
}

func (s *materialService) HomeMaterialData(c context.Context, status int) (int64, error) {
	count, err := s.materialRepository.HomeMaterialData(c, status)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *materialService) MaterialUploadTrend(c context.Context) (*v1.ReqListListInt, error) {
	resp := &v1.ReqListListInt{
		ListListInt: make([]v1.ReqListInt, 0, 2),
	}

	List1, err := s.materialRepository.MaterialUploadTrend(c, 1)
	if err != nil {
		return nil, err
	}

	List2, err := s.materialRepository.MaterialUploadTrend(c, 0)
	if err != nil {
		return nil, err
	}

	resp.ListListInt = append(resp.ListListInt, v1.ReqListInt{
		ListInt: List1,
		Status:  "视频",
	})

	resp.ListListInt = append(resp.ListListInt, v1.ReqListInt{
		ListInt: List2,
		Status:  "图片",
	})
	return resp, nil
}

func (s *materialService) UpAuditMaterial(c context.Context, req *v1.UpAuditMaterialReq) error {
	err := s.materialRepository.UpAuditMaterial(c, req)
	if err != nil {
		return v1.ErrUpdataStatus
	}
	return nil
}

func (s *materialService) GetPcCountData(c context.Context, req *v1.PcData) (int64, error) {
	Count, err := s.materialRepository.GetPcCountData(c, req)
	if err != nil {
		return 0, v1.ErrGetPcDataCount
	}
	return Count, nil
}

func (s *materialService) GetPcUrlDataList(c context.Context, req *v1.PcData) (*v1.ResPcDataList, error) {
	materials, err := s.materialRepository.GetPcUrlDataList(c, req)
	if err != nil {
		return nil, v1.ErrGetPcUrlData
	}

	resp := &v1.ResPcDataList{
		ResPcData: make([]v1.ResPcData, 0, len(materials)),
	}

	for _, material := range materials {
		resp.ResPcData = append(resp.ResPcData, v1.ResPcData{
			Id:material.ID,
			Name:material.Name,
			Url:material.ImageUrl,
			Label: material.LabelId,
			CategoryNameList: "暂时无法获取",
		})
	}
	return resp, nil
}

func (s *materialService) DeletePcMaterial(c context.Context, req *v1.PcData) error {
	err := s.materialRepository.DeletePcMaterial(c, req.Id)
	if err != nil {
		return v1.ErrDeletePcMaterial
	}

	return nil
}