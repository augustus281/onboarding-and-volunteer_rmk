package dto

// SignInRequestDTO is used to capture the data required for a user sign-in request.
// SignInRequestDTO được sử dụng để thu thập dữ liệu cần thiết cho yêu cầu đăng nhập của người dùng.
type SignInRequestDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SignInResponseDTO is used to send back a response after a successful sign-in.
// SignInResponseDTO được sử dụng để gửi lại phản hồi sau khi đăng nhập thành công.
type SignInResponseDTO struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	UserID    uint   `json:"user_id"`
	Token     string `json:"token"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// SignUpRequestDTO is used to capture the data required for a user sign-up request.
// SignUpRequestDTO được sử dụng để thu thập dữ liệu cần thiết cho yêu cầu đăng ký của người dùng.
type SignUpRequestDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	UserID   uint   `json:"user_id" binding:"required"`
}

// SignUpResponseDTO is used to send back a response after a successful sign-up.
// SignUpResponseDTO được sử dụng để gửi lại phản hồi sau khi đăng ký thành công.
type SignUpResponseDTO struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	UserID    uint   `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
