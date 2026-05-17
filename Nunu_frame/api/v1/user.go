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
type DeleteUserRequest struct {
	UserID uint `json:"user_id" binding:"required" example:"1"`
}
type DeleteUserResponse struct {
	UserID    uint   `json:"user_id"`
	DeletedAt string `json:"deleted_at"`
}

type UpdateUserRequest struct {
	UserID   uint   `json:"user_id" binding:"required" example:"1"`
	Account  string `json:"account" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"123456"`
	Email    string `json:"email" binding:"required" example:"admin@example.com"`
	Phone    string `json:"phone" binding:"required" example:"1234567890"`
	RealName string `json:"real_name" binding:"required" example:"Admin User"`
}
type UpdateUserResponse struct {
	UserID    uint   `json:"user_id"`
	Account   string `json:"account"`
	UpdatedAt string `json:"updated_at"`
}
type ResetUserPasswordRequest struct {
	UserID   uint   `json:"user_id" binding:"required" example:"1"`
	Password string `json:"password" binding:"required" example:"123456"`
}
type ResetUserPasswordResponse struct {
	UserID    uint   `json:"user_id"`
	UpdatedAt string `json:"updated_at"`
}
