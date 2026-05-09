package v1

// 实现查询分类列表接口结构
// 对应schooldemo数据库中的classification 表
type GetClassReq struct {
	Id        int    `json:"id"`
	className string `json:"className"`
}

type GetClassSReq struct {
	Id      int    `json:"id"`
	name    string `json:"name"`
	classId int    `json:"classId"`
}
