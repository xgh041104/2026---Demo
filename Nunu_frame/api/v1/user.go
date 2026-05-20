package v1

type LoginRequest struct {
	Account  string `json:"account" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"123456"`
}
type LoginResponse struct {
	Account string `json:"account"`
	Token   string `json:"token"`
}

type CreateUserRequest struct {
	Account  string `json:"account" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"123456"`
}
type CreateUserResponse struct {
	UserID    uint   `json:"user_id"`
	Account   string `json:"account"`
	CreatedAt string `json:"created_at"`
}

type UpdateUserRequest struct {
	Account    string `json:"account" example:"admin"`
	Password   string `json:"password"  example:"123456"`
	RealName   string `json:"real_name"  example:"admin"`
	Phone      string `json:"phone"  example:"1234567890"`
	Email      string `json:"email"  example:"admin@example.com"`
	Type       string `json:"type"  example:"user"`
	IsDisabled bool   `json:"is_disabled"  example:"false"`
}

type FindUserResponse struct {
	Account    string `json:"account" example:"admin"`
	RealName   string `json:"real_name"  example:"admin"`
	Phone      string `json:"phone"  example:"1234567890"`
	Email      string `json:"email"  example:"admin@example.com"`
	Type       string `json:"type"  example:"user"`
	IsDisabled bool   `json:"is_disabled"  example:"false"`
}
type UserListResponse struct {
	Total int64              `json:"total"`
	List  []FindUserResponse `json:"list"`
}
type FindUserRequest struct {
	Account  string `json:"account" example:"admin"`
	RealName string `json:"real_name" example:"admin"`
	Phone    string `json:"phone" example:"1234567890"`
	Page     int    `json:"page" example:"1"`
	PageSize int    `json:"page_size" example:"10"`
}
