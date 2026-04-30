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

	//label
	ErrFindLabel         = newError(1004, "标签查找失败")
	ErrCreateLabel       = newError(1005, "标签创建失败")
	ErrLabelAlreadyExist = newError(1006, "标签已存在")

	//category
	ErrFindCategory = newError(1005, "分类查找失败")

	//subCategory
	ErrFindSubCategory = newError(1006, "子分类查找失败")
)
