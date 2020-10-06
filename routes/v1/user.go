package v1

import (
	c "github.com/Team-IF/TeamWork_Backend/controllers/v1/auth"
	m "github.com/Team-IF/TeamWork_Backend/middlewares"
	req "github.com/Team-IF/TeamWork_Backend/models/req"
	"github.com/gin-gonic/gin"
)

func SetUserRoutes(r *gin.RouterGroup) {

	r.POST("/signin", m.VerifyRequest(&req.UserSignIn{}), c.SignIn)
	r.POST("/signup", m.VerifyRequest(&req.UserSignUp{}), c.SignUp)

	// r.POST("/resetpasswordverify", m.VerifyRequest(&req.ResetPasswordGetCode{}), c.ResetPasswordGetCode)

	// r.POST("/updatepassword", m.VerifyRequest(&req.ResetPassword{}), m.CheckAuth(), c.UpdatePassword)

	// r.POST("/emailverify", c.Verify)

	// r.GET("/validate", m.CheckAuth(), c.Validate)
	// r.GET("/refresh", c.Refresh)

	// r.POST("/resetpassword", m.VerifyRequest(&req.ResetPasswordWithCode{}), c.ResetPasswordWithCode)
}
