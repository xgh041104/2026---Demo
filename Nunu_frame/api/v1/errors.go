package v1

var (
	// common errors
	ErrSuccess             = newError(200, "ok")
	ErrBadRequest          = newError(400, "Bad Request")
	ErrUnauthorized        = newError(401, "Unauthorized")
	ErrNotFound            = newError(404, "Not Found")
	ErrInternalServerError = newError(500, "Internal Server Error")
	ErrPasswordNotNil      = newError(1000, "密码不能为空")
	ErrFindUser            = newError(1002, "用户不存在")
	ErrUserAlreadyExist    = newError(1003, "用户已存在")
	ErrPasswordEncrypt     = newError(1004, "密码加密错误")
	ErrNotFindUser         = newError(1005, "用户不存在")

	// more biz errors
	ErrEmailAlreadyUse = newError(1001, "The email is already in use.")

	//参数绑定错误
	ErrBind = newError(1002, "Bind Error")

	//user
	ErrPasswordIncorrect = newError(1003, "密码不匹配")
	ErrUserNotFound      = newError(1004, "用户不存在")
	ErrUserFailFind      = newError(1005, "用户查询失败")
	ErrUserNotDisabled   = newError(1006, "用户未禁用,不能删除")
	ErrUpdateUser        = newError(1007, "更新用户失败")

	//label
	ErrFindLabel         = newError(1004, "标签查找失败")
	ErrCreateLabel       = newError(1005, "标签创建失败")
	ErrLabelAlreadyExist = newError(1006, "标签已存在")
	ErrDeleteLabel       = newError(1007, "标签删除失败")

	//category
	ErrFindCategory     = newError(1005, "分类查找失败")
	ErrCategoryExists   = newError(1006, "分类已存在")
	ErrCategoryNotExist = newError(1007, "分类不存在")
	ErrDeleteCategory   = newError(1008, "分类删除失败")

	//subCategory
	ErrFindSubCategory         = newError(1006, "子分类查找失败")
	ErrDeleteSubCategory       = newError(1007, "子分类删除失败")
	ErrSubCategoryAlreadyExist = newError(1008, "子分类已存在")
	ErrSubCategoryNotExist     = newError(1009, "子分类不存在")

	// material
	ErrCreateUser          = newError(1009, "创建用户失败")
	ErrDeleteUser          = newError(1010, "删除用户失败")
	ErrReset               = newError(1011, "重置密码失败")
	ErrFindAllAccount      = newError(1012, "查找所有账号失败")
	ErrFindLikeAccount     = newError(1013, "查找相似账号失败")
	ErrPicUserId           = newError(1014, "用户ID不能为空")
	ErrPicUrl              = newError(1015, "URL不能为空")
	ErrGetMaterialList     = newError(1016, "获取素材库列表失败")
	ErrGetMaterialLabelStr = newError(1017, "获取标签表组失败")
	ErrUpdataStatus        = newError(1018, "更新上传状态失败")
	ErrGetPcDataCount      = newError(1019, "获取数据库个人数据数量失败")
	ErrGetPcUrlData        = newError(1020, "获取个人数据Url失败")
	ErrDeletePcMaterial    = newError(1021, "删除素材失败")
	ErrFileExt             = newError(1022, "文件后缀错误")
)
