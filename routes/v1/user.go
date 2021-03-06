package v1

import (
	c "github.com/Team-IF/TeamWork_Backend/controllers/v1/auth"
	m "github.com/Team-IF/TeamWork_Backend/middlewares"
	req "github.com/Team-IF/TeamWork_Backend/models/req"
	"github.com/gin-gonic/gin"
)

func setUserRoutes(r *gin.RouterGroup) {

	r.GET("/@me", m.CheckAuth(), c.GetProile)
	r.PATCH("/@me", m.CheckAuth(), m.VerifyRequest(&req.UserUpdateProfile{}), c.UpdateProfile)
	r.PATCH("/@me/password", m.CheckAuth(), m.VerifyRequest(&req.UserUpdatePassword{}), c.UpdatePassword)

	r.GET("/signin", m.VerifyQuery(&req.UserSignIn{}), c.SignIn)
	r.POST("/signup", m.VerifyRequest(&req.UserSignUp{}), c.SignUp)

	r.POST("/emailverify", m.VerifyRequest(&req.UserVerifyEmail{}), c.VerifyEmail)

	r.GET("/validate", m.CheckAuth(), c.Validate)
	r.GET("/refresh", c.Refresh)

	r.POST("/resetpasswordverify", m.VerifyRequest(&req.UserResetPasswordVerify{}), c.ResetPasswordVerify)
	r.POST("/resetpassword", m.VerifyRequest(&req.UserResetPassword{}), c.ResetPassword)
}
