package v1

type PageNumAndPageSizeRequest struct {
	PageNum  int `form:"pageNum"`
	PageSize int `form:"pageSize"`
}

type PageNumSizeAndStatus struct {
	PageNum  int `form:"pageNum"`
	PageSize int `form:"pageSize"`
	Status   int `form:"status"`
}

// 审核中心-审核素材Res
type ReqAuditingMaterial struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	CreatorName string `json:"creator_name"`
	CreatedAt   string `json:"created_at"`
	ImageUrl    string `json:"image_url"`
}

type ReqAuditingMaterialListResponse struct {
	ReqMaterialList []ReqAuditingMaterial `json:"list"`
}

// 审核中心
type UpAuditMaterialReq struct {
	Id     uint   `json:"id"`
	Status int    `json:"status"`
	Remark string `json:"remark"`
}

// 个人中心
type PcData struct {
	Id     uint `json:"id" form:"id"`
	Status int  `json:"status" form:"status"`
}

type ResPcData struct {
	Id               uint   `json:"id"`
	Name             string `json:"name"`
	Url              string `json:"Url"`
	Label            string `json:"label"`
	CategoryNameList string `json:"CategoryList"`
}

type ResPcDataList struct {
	ResPcData []ResPcData `json:"list"`
}

// 素材页面-Get素材列表
type ReqMaterialRes struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	CreatedAt  string `json:"created_at"`
	ImageUrl   string `json:"image_url"`
	CategoryId int    `json:"category_id"`
	LabelName  string `json:"label_name"`
}

type ReqMaterialListRes struct {
	ReqMaterialResList []ReqMaterialRes `json:"list"`
}

// 首页-上传趋势
type ReqListInt struct {
	ListInt []int64 `json:"upload_count"`
	Status  string  `json:"status"`
}

type ReqListListInt struct {
	ListListInt []ReqListInt `json:"List_upload_count"`
}


