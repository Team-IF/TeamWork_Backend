package auth

import (
	"fmt"

	"github.com/Team-IF/TeamWork_Backend/db"
	"github.com/Team-IF/TeamWork_Backend/models/req"
	"github.com/Team-IF/TeamWork_Backend/models/res"
	"github.com/Team-IF/TeamWork_Backend/utils"
	resutil "github.com/Team-IF/TeamWork_Backend/utils/res"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignIn(c *gin.Context) {
	r := resutil.New(c)
	query := c.MustGet("query").(*req.UserSignIn)
	data, err := db.SignIn(query.ID, query.Password)
	if err != nil {
		if err == gorm.ErrRecordNotFound || err == db.ErrPasswordNotMatch {
			r.SendError(resutil.ERR_AUTH, "Username or Password is incorrect")
		} else {
			r.SendError(resutil.ERR_SERVER, "Error while signin")
		}
		fmt.Println(err)
		return
	}

	token, err := utils.GetJwtToken(data.ID)
	if err != nil {
		r.SendError(resutil.ERR_SERVER, "Error while making token")
		return
	}

	r.Response(&res.User{
		Name:          data.Name,
		DisplayName:   data.DisplayName,
		Avatar:        data.Avatar,
		Email:         data.Email,
		EmailVerified: data.EmailVerified,
		Token:         token,
	})
}
