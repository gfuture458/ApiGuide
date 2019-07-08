package controllers

import (
	"ApiGuide/models"
	"ApiGuide/utils"
	"ApiGuide/validations"
	"encoding/json"
	"strconv"
)

type InterfaceController struct {
	BaseController
}

func (inter *InterfaceController) AddInterface() {
	var targetData models.Interface
	_ = json.Unmarshal(inter.GetData(), &targetData)
	errList := validations.CheckInterface(&targetData)
	if len(errList) > 0 {
		inter.jsonResult(utils.FalseReturn(-1, "错误的请求", errList))
	}
	exist := models.GetInterfaceById(targetData.Id)
	if exist {
		if _, er := models.UpdateInterface(&targetData); er == nil {
			inter.jsonResult(utils.TrueReturn("接口更新成功", targetData.Id))
		} else {
			inter.jsonResult(utils.FalseReturn(-1, "接口更新失败", er))
		}
	} else {
		targetData.IsActive = true
		if id, err := models.AddInterface(&targetData); err == nil {
			inter.jsonResult(utils.TrueReturn("接口添加成功", id))
		} else {
			inter.jsonResult(utils.FalseReturn(-1, "接口添加失败", err))
		}
	}
}

func (inter *InterfaceController) GetInterface() {
	mid := inter.GetString("mid")
	targer, _ := strconv.Atoi(mid)
	if interfaces, err := models.GetInterface(targer); err == nil {
		inter.jsonResult(utils.TrueReturn("请求成功", interfaces))
	} else {
		inter.jsonResult(utils.FalseReturn(-1, "请求失败", err))
	}
}
