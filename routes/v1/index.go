package v1

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.RouterGroup) {
	setUserRoutes(r.Group("user"))
	setProjectRoutes(r.Group("project"))
}
