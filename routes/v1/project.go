package v1

import (
	"github.com/gin-gonic/gin"

	c "github.com/Team-IF/TeamWork_Backend/controllers/v1/project"
	m "github.com/Team-IF/TeamWork_Backend/middlewares"
	req "github.com/Team-IF/TeamWork_Backend/models/req"
)

func setProjectRoutes(r *gin.RouterGroup) {
	r.POST("/", m.VerifyRequest(&req.ProjectCreate{}), c.CreateProject)
	r.PATCH("/", m.VerifyRequest(&req.ProjectUpdate{}), c.UpdateProject)
}
