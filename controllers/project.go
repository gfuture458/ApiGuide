package controllers

import (
	"ApiGuide/models"
	"ApiGuide/utils"
	"ApiGuide/validations"
	"encoding/json"
	"fmt"
	"strconv"
)

type ProjectController struct {
	BaseController
}

func (pro *ProjectController) AdPro() {
	var project models.Project
	_ = json.Unmarshal(pro.GetData(), &project)
	errList := validations.ProjectAddCheck(&project)
	if len(errList) > 0 {
		pro.jsonResult(utils.FalseReturn(-1, "参数错误", errList))
	}
	fmt.Println(project.Name)
	target, _ := models.GetProjectByName(project.Name)
	if target != nil {
		pro.jsonResult(utils.FalseReturn(-1, "项目已存在", ""))
	}
	project.User = pro.user
	project.IsActive = true
	_, err := models.AddProject(&project)
	if err != nil {
		pro.jsonResult(utils.FalseReturn(-1, "项目添加失败", err))
	} else {
		pro.jsonResult(utils.TrueReturn("项目添加成功", ""))
	}
}

func (pro *ProjectController) GetAllPjc() {
	ProList, err := models.GetAllProject(pro.user.Id)
	if err != nil {
		pro.jsonResult(utils.FalseReturn(-1, "查询错误", err))
	}
	pro.jsonResult(utils.TrueReturn("请求成功", ProList))
}

func (pro *ProjectController) DelPjt() {
	pid := pro.GetString("id")
	//fmt.Println(reflect.TypeOf(pid))
	delId, _ := strconv.Atoi(pid)
	result := models.DeleteProject(delId)
	if result {
		pro.jsonResult(utils.TrueReturn("删除成功", ""))
	}
	pro.jsonResult(utils.FalseReturn(-1, "删除失败", ""))
}
