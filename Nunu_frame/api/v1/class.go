package v1

// 实现查询分类列表接口结构
// 对应schooldemo数据库中的classification 表
type GetClassReq struct {
	Id        int    `json:"id"`
	ClassName string `json:"className"`
}

type GetClassSReq struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	ClassId int    `json:"classId"`
}
