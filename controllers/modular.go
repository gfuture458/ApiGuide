package controllers

import (
	"ApiGuide/models"
	"ApiGuide/utils"
	"ApiGuide/validations"
	"encoding/json"
	"strconv"
)

type ModularController struct {
	BaseController
}

func (modular *ModularController) GetAllModular() {
	pid := modular.GetString("pid")
	p, _ := strconv.Atoi(pid)
	list, err := models.GetModulars(p)
	if err != nil {
		modular.jsonResult(utils.FalseReturn(-1, "系统错误", err))
	}
	modular.jsonResult(utils.TrueReturn("请求成功", list))
}

func (modular *ModularController) AddModular() {
	var targetModular models.Modular
	_ = json.Unmarshal(modular.GetData(), &targetModular)
	errList := validations.CheckModular(&targetModular)
	if len(errList) != 0 {
		modular.jsonResult(utils.FalseReturn(-1, "参数错误", errList))
	}
	targetModular.IsActive = true
	result, err := models.AddModular(&targetModular)
	if err != nil {
		modular.jsonResult(utils.FalseReturn(-1, "添加失败", err))
	}
	modular.jsonResult(utils.TrueReturn("添加成功", result))
}
