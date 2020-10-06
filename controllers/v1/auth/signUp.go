package auth

import (
	"github.com/Team-IF/TeamWork_Backend/db"
	"github.com/Team-IF/TeamWork_Backend/models/req"
	"github.com/Team-IF/TeamWork_Backend/models/res"
	"github.com/Team-IF/TeamWork_Backend/utils"
	resutil "github.com/Team-IF/TeamWork_Backend/utils/res"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignUp(c *gin.Context) {
	r := resutil.New(c)
	body := c.MustGet("body").(*req.UserSignUp)
	verifyCode := utils.CreateRandomString(8)
	userID, err := db.CreateUser(body.Name, body.DisplayName, body.Avatar, body.Email, body.Password, verifyCode)
	if err != nil {
		r.SendError(resutil.ERR_DUPLICATE, "Duplicate")
		return
	}

	if err := utils.SendVefiryMail(verifyCode, body.Email); err != nil {
		r.SendError(resutil.ERR_SERVER, "Error Send Verify Mail")
		return
	}

	token, err := utils.GetJwtToken(userID)
	if err != nil {
		r.SendError(resutil.ERR_SERVER, "Error Get JWT Token")
		return
	}
	r.Response(res.Token{Token: token})
}

func VerifyEmail(c *gin.Context) {
	r := resutil.New(c)
	body := c.MustGet("body").(*req.UserVerifyEmail)
	if err := db.VerifyEmail(body.Email, body.VerifyCode); err != nil {
		if err == gorm.ErrRecordNotFound {
			r.SendError(resutil.ERR_BAD_REQUEST, "Code or Email is not match")
		} else if err == db.ErrAlreadyVerified {
			r.SendError(resutil.ERR_BAD_REQUEST, "This Account is already Verified")
		} else if err == db.ErrExpired {
			r.SendError(resutil.ERR_BAD_REQUEST, "verify code expired")
		} else {
			r.SendError(resutil.ERR_SERVER, "Error while select")
		}
		return
	}
	r.Response(resutil.R{})
}
