package v1

type LoginRequest struct {
	Account  string `json:"account" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"123456"`
}
type LoginResponse struct {
	Account string `json:"account"`
	Token   string `json:"token"`
}
