package req

type UserSignUp struct {
	UserID   string `json:"userid" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Avatar   string `json:"avatar"`
	Password string `json:"password" binding:"required"`
}

type UserSignIn struct {
	ID       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdatePassword struct {
	Origin string `json:"origin" binding:"required"`
	New    string `json:"new" binding:"required"`
}

type UserResetPasswordVerify struct {
	Email string `json:"email" binding:"required"`
}

type UserResetPassword struct {
	Verify   string `json:"verify" binding:"required"`
	Password string `json:"password" binding:"required"`
}
