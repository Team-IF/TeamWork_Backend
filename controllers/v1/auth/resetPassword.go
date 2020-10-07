package auth

import (
	"github.com/Team-IF/TeamWork_Backend/db"
	"github.com/Team-IF/TeamWork_Backend/models/req"
	"github.com/Team-IF/TeamWork_Backend/utils"
	resutil "github.com/Team-IF/TeamWork_Backend/utils/res"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ResetPasswordVerify(c *gin.Context) {
	r := resutil.New(c)
	body := c.MustGet("body").(*req.UserResetPasswordVerify)

	verifyCode := utils.CreateRandomString(8)
	if err := db.PasswordVerifyCode(body.Email, verifyCode); err != nil {
		r.SendError(resutil.ERR_SERVER, "error while set password")
		return
	}

	if err := utils.SendResetPasswordMail(verifyCode, body.Email); err != nil {
		r.SendError(resutil.ERR_SERVER, "error while sending email")
		return
	}
	r.Response(&resutil.R{})
}

func ResetPassword(c *gin.Context) {
	r := resutil.New(c)
	body := c.MustGet("body").(*req.UserResetPassword)

	if err := db.UpdatePasswordWithCode(body.Email, body.Password, body.Verify); err != nil {
		if err == gorm.ErrRecordNotFound {
			r.SendError(resutil.ERR_NOT_MATCH, "Code or Email is not match")
		} else if err == db.ErrExpired {
			r.SendError(resutil.ERR_EXPIRED, "verify code expired")
		} else {
			r.SendError(resutil.ERR_SERVER, "Error while select")
		}
		return
	}
	r.Response(resutil.R{})
}
