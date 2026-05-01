package v1

type CreateCategoryReq struct {
	Name   string `json:"name"`
	Remark string `json:"remark"`
	Sort   uint   `json:"sort"`
	Status uint   `json:"status"`
}
