package auth

import (
	resutil "github.com/Team-IF/TeamWork_Backend/utils/res"
	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context) {
	resutil.New(c).Response(resutil.R{})
}

func Refresh(c *gin.Context) {
	// r := resutil.New(c)
}
