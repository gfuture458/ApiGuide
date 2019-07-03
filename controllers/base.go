package controllers

import (
	"ApiGuide/models"
	"ApiGuide/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	user           *models.User
	userId         int
	userName       string
}

func (this *BaseController) Prepare() {
	//  登陆认证
	var noAuth = [...]string{
		"/v1/user/register",
		"/v1/user/login",
	}
	reqUrl := this.Ctx.Request.RequestURI
	auth := true
	for _, a := range noAuth {
		if a == reqUrl {
			auth = false
		}
	}
	if auth {
		this.auth()
	}
}

func (this *BaseController) auth() {
	req := this.Ctx.Request.Header.Get("Authorization")
	if req == "" {
		this.jsonResult(utils.FalseReturn(600, "认证失败", ""))
	}
	token := strings.Split(req, " ")
	addr := this.getClientAddr()
	reqUser, err := utils.ParseToken(token[1], addr)
	if err != nil {
		this.jsonResult(utils.FalseReturn(600, "认证失败", ""))
	}
	if reqUser.Id == 0 {
		this.jsonResult(utils.FalseReturn(600, "认证失败", ""))
	}
	this.user = reqUser
}

func (this *BaseController) getJson() map[string]interface{} {
	var demo map[string]interface{}
	res := this.Ctx.Request.Body
	formValye, _ := ioutil.ReadAll(res)
	r := json.Unmarshal(formValye, &demo)
	if r != nil {
		out := make(map[string]interface{})
		out["status"] = "系统异常"
		out["code"] = -1
		this.jsonResult(out)
	}
	return demo
}

// 返回json数据
func (this *BaseController) jsonResult(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) getClientAddr() string {
	ip := strings.Split(this.Ctx.Request.RemoteAddr, ":")[0]
	return ip
}

func (this *BaseController) GetData() []uint8 {
	reqData, _ := ioutil.ReadAll(this.Ctx.Request.Body)
	return reqData
}
