package project

import (
	"github.com/Team-IF/TeamWork_Backend/db"
	"github.com/Team-IF/TeamWork_Backend/models/req"
	resutil "github.com/Team-IF/TeamWork_Backend/utils/res"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProject(c *gin.Context) {
	r := resutil.New(c)
	body := c.MustGet("body").(*req.ProjectCreate)
	userID := c.MustGet("uerID").(uint)
	projectID, err := db.CreateProject(userID, body.Name, body.Description, body.Password)
	if err != nil {
		r.SendError(resutil.ERR_SERVER, "Error while writing data to db")
		return
	}
	r.Response(resutil.R{
		"ID": projectID,
	})
}

func UpdateProject(c *gin.Context) {
	r := resutil.New(c)
	body := c.MustGet("body").(*req.ProjectUpdate)
	userID := c.MustGet("uerID").(uint)
	err := db.UpdateProject(userID, body.ID, body.Name, body.Description, body.Password)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			r.SendError(resutil.ERR_NOT_FOUND, "project not found")
		} else if err == db.ErrNoPermission {
			r.SendError(resutil.ERR_NO_PERMISSION, "user doesn't have permission")
		} else {
			r.SendError(resutil.ERR_SERVER, "Error while writing data to db")
		}
		return
	}
	r.Response(resutil.R{})
}
