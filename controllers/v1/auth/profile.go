package auth

import (
	"github.com/Team-IF/TeamWork_Backend/db"
	dbmodels "github.com/Team-IF/TeamWork_Backend/models/db"
	"github.com/Team-IF/TeamWork_Backend/models/req"
	"github.com/Team-IF/TeamWork_Backend/models/res"
	"github.com/Team-IF/TeamWork_Backend/utils"
	resutil "github.com/Team-IF/TeamWork_Backend/utils/res"
	"github.com/gin-gonic/gin"
)

func GetProile(c *gin.Context) {
	r := resutil.New(c)
	user := c.MustGet("user").(*dbmodels.User)

	r.Response(&res.User{
		Name:          user.Name,
		DisplayName:   user.DisplayName,
		Avatar:        user.Avatar,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
	})
}

func UpdateProfile(c *gin.Context) {
	r := resutil.New(c)
	user := c.MustGet("user").(*dbmodels.User)
	body := c.MustGet("body").(*req.UserUpdateProfile)

	if err := db.UpdateProfile(user.ID, body.Name, body.DisplayName, body.Avatar, body.Email); err != nil {
		r.SendError(resutil.ERR_SERVER, "error while write db")
		return
	}

	r.Response(&resutil.R{})
}

func UpdatePassword(c *gin.Context) {
	r := resutil.New(c)
	user := c.MustGet("user").(*dbmodels.User)
	body := c.MustGet("body").(*req.UserUpdatePassword)

	if !utils.CheckPassword(body.Origin, user.Password) {
		r.SendError(resutil.ERR_AUTH, "origin password not match")
		return
	}

	if err := db.UpdatePassword(user.ID, body.New); err != nil {
		r.SendError(resutil.ERR_SERVER, "error while write db")
		return
	}

	r.Response(&resutil.R{})
}
