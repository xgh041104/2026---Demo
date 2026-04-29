package v1

type LoginRequest struct {
	Account  string `json:"account" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"123456"`
}
type LoginResponse struct {
	Account string `json:"account"`
	Token   string `json:"token"`
}

type CreateUserReq struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 修改接口
type UpdateUserReq struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
	RealName string `json:"real_name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// 删除接口
type DeleteUserReq struct {
	Account    string `gorm:"type:varchar(50);not null;unique;comment:登录账号" json:"account"`
	Password   string `gorm:"type:varchar(100);not null;comment:登录密码" json:"password"`
	RealName   string `gorm:"type:varchar(50);not null;comment:真实姓名" json:"real_name"`
	Phone      string `gorm:"type:varchar(20);comment:手机号" json:"phone"`
	Email      string `gorm:"type:varchar(100);comment:邮箱" json:"email"`
	Type       string `gorm:"type:varchar(20);not null;comment:用户角色 admin/root/user" json:"type"`
	IsDisabled int8   `gorm:"default:0;comment:是否禁用 0启用 1禁用" json:"is_disabled"`
}

type ResetUserPasswordReq struct {
	Account  string `gorm:"type:varchar(50);not null;unique;comment:登录账号" json:"account"`
	Password string `gorm:"type:varchar(100);not null;comment:登录密码" json:"password"`
}
