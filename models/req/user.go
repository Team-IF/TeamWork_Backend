package req

type UserSignUp struct {
	Name        string `json:"userid" binding:"required"`
	DisplayName string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Avatar      string `json:"avatar"`
	Password    string `json:"password" binding:"required"`
}

type UserUpdateProfile struct {
	Name        string `json:"userid" binding:"required"`
	DisplayName string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Avatar      string `json:"avatar"`
}

type UserVerifyEmail struct {
	Email      string `json:"email" binding:"required"`
	VerifyCode string `json:"verify_code" binding:"required"`
}

type UserSignIn struct {
	ID       string `form:"userid"`
	Password string `form:"password"`
}

type UserUpdatePassword struct {
	Origin string `json:"origin" binding:"required"`
	New    string `json:"new" binding:"required"`
}

type UserResetPasswordVerify struct {
	Email string `json:"email" binding:"required"`
}

type UserResetPassword struct {
	Email    string `json:"email" binding:"required"`
	Verify   string `json:"verify" binding:"required"`
	Password string `json:"password" binding:"required"`
}
