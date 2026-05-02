package v1

type CreateCategoryReq struct {
	Name   string `json:"name"`
	Remark string `json:"remark"`
	Sort   uint   `json:"sort"`
	Status uint   `json:"status"`
}

type DeleteCategoryReq struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
	Sort   uint   `json:"sort"`
	Status uint   `json:"status"`

	SubCategories []*SubCategory `json:"sub_categories"`
}
