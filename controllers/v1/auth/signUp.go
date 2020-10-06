package auth

import (
	"github.com/Team-IF/TeamWork_Backend/db"
	"github.com/Team-IF/TeamWork_Backend/models/req"
	"github.com/Team-IF/TeamWork_Backend/models/res"
	"github.com/Team-IF/TeamWork_Backend/utils"
	resutil "github.com/Team-IF/TeamWork_Backend/utils/res"
	"github.com/gin-gonic/gin"
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
	// r := resutil.New(c)
}
