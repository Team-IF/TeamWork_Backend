package db

import (
	dbmodels "github.com/Team-IF/TeamWork_Backend/models/db"
	"github.com/Team-IF/TeamWork_Backend/utils"
)

func CreateProject(id uint, projectName, projectDescription, projectPassword string) (uint, error) {
	project := &dbmodels.Project{
		Name:         projectName,
		ProjectOwner: id,
		Description:  projectDescription,
		Password:     projectPassword,
	}
	err := utils.GetDB().Create(&project).Error
	if err != nil {
		return 0, err
	}
	return project.ID, err
}

func GetProject(projectID uint) (*dbmodels.Project, error) {
	var project dbmodels.Project
	err := utils.GetDB().Find(&project, projectID).Error
	return &project, err
}

func CheckProjectOwner(userID, projectID uint) (bool, error) {
	var project dbmodels.Project
	err := utils.GetDB().Select("project_owner").Find(&project, projectID).Error
	if err != nil {
		return false, err
	}
	return project.ProjectOwner == userID, err
}

func UpdateProject(userID, projectID uint, projectName, projectDescription, projectPassword string) error {
	isOwner, err := CheckProjectOwner(userID, projectID)
	if err != nil {
		return err
	} else if !isOwner {
		return ErrNoPermission
	}
	err = utils.GetDB().Model(&dbmodels.Project{}).Where("id = ?", projectID).Updates(&dbmodels.Project{
		Name:        projectName,
		Description: projectDescription,
		Password:    projectPassword,
	}).Error

	return err
}

func UserProjectCount(id uint) (int, error) {
	var count int64
	err := utils.GetDB().Where("user_id = ?", id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), err
}

func UserProjectList(id uint, limit, offset int) (*[]dbmodels.Project, error) {
	var projectList []dbmodels.ProjectMemeber
	err := utils.GetDB().Where("user_id = ?", id).Limit(limit).Offset(offset).Find(&projectList).Error
	if err != nil {
		return nil, err
	}
	var projectDetailList []dbmodels.Project
	for _, item := range projectList {
		var projectDetail dbmodels.Project
		err := utils.GetDB().Where("project_id = ", item.ProjectID).Find(&projectDetail).Error
		if err != nil {
			return nil, err
		}
		projectDetailList = append(projectDetailList, projectDetail)
	}
	return &projectDetailList, nil
}
