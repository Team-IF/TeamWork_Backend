package middlewares

import (
	resutil "github.com/Team-IF/TeamWork_Backend/utils/res"
	"github.com/gin-gonic/gin"
)

func VerifyRequest(data interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := resutil.New(c)
		if err := c.ShouldBindJSON(data); err != nil {
			r.SendError(resutil.ERR_BAD_REQUEST, err.Error())
			return
		}
		c.Set("body", data)
	}
}

func VerifyQuery(data interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := resutil.New(c)
		if err := c.ShouldBindQuery(data); err != nil {
			r.SendError(resutil.ERR_BAD_REQUEST, err.Error())
			return
		}
		c.Set("query", data)
	}
}
