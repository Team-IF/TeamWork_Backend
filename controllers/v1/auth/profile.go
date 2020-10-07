package auth

import (
	dbmodels "github.com/Team-IF/TeamWork_Backend/models/db"
	"github.com/Team-IF/TeamWork_Backend/models/res"
	resutil "github.com/Team-IF/TeamWork_Backend/utils/res"
	"github.com/gin-gonic/gin"
	// resutil "github.com/Team-IF/TeamWork_Backend/utils/res"
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
	// r := resutil.New(c)
}

func UpdatePassword(c *gin.Context) {
	r := resutil.New(c)
	user := c.MustGet("user").(*dbmodels.User)

	r.Response(&resutil.R{})
}
