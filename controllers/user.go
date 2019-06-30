package controllers

import (
	"ApiGuide/models"
	"ApiGuide/utils"
	"ApiGuide/validations"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

// Operations about Users
type UserController struct {
	BaseController
}

// @Title 注册用户
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Register() {
	var user models.User
	_ = json.Unmarshal(u.GetData(), &user)
	list := validations.RegisterCheck(&user)
	if len(list) > 0 {
		u.jsonResult(utils.FalseReturn(-1, "参数错误", list))
	}
	user.Password = base64.StdEncoding.EncodeToString([]byte(user.Password))
	user.Login = time.Now()
	exist, _ := models.GetUserByName(user.Username)
	if exist != nil {
		u.jsonResult(utils.FalseReturn(-1, "账号已存在", ""))
	}
	_, err := models.AddUser(&user)
	if err != nil {
		fmt.Println(err)
		u.jsonResult(utils.FalseReturn(-1, "账号注册失败", ""))
	} else {
		u.jsonResult(utils.TrueReturn("注册成功", ""))
	}
}

// @Title 登陆
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	var user models.User
	_ = json.Unmarshal(u.GetData(), &user)
	error_list := validations.LoginCheck(&user)
	if len(error_list) > 0 {
		u.jsonResult(utils.FalseReturn(-1, "参数错误", error_list))
	}
	target_user, err := models.GetUserByName(user.Username)
	if err != nil {
		u.jsonResult(utils.FalseReturn(-1, "账号不存在", err.Error()))
	}
	if target_user.Password == user.Password {
		target_user.Login = time.Now()
		_ = models.UpdateUser(target_user)
		token, _ := utils.CreateToken(target_user, u.getClientAddr())
		target := make(map[string]interface{})
		target["code"] = 200
		target["msg"] = "登陆成功"
		target["token"] = token
		u.jsonResult(target)
	} else {
		u.jsonResult(utils.FalseReturn(-1, "密码错误", ""))
	}
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users, err := models.GetAllUsers()
	if err != nil {
		u.jsonResult(utils.FalseReturn(-1, "请求错误", ""))
	}
	u.jsonResult(utils.TrueReturn("请求成功", users))
}

//@Title Get
//@Description get user by uid
//@Param	uid		path 	string	true		"The key for staticblock"
//@Success 200 {object} models.User
//@Failure 403 :uid is empty
//@router /:uid [get]
func (u *UserController) GetUserInfo() {
	target_user, _ := models.GetUserByName(u.user.Username)
	u.jsonResult(utils.TrueReturn("请求成功", target_user))
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
//func (u *UserController) Put() {
//	uid := u.GetString(":uid")
//	if uid != "" {
//		var user models.User
//		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
//		uu, err := models.UpdateUser(uid, &user)
//		if err != nil {
//			u.Data["json"] = err.Error()
//		} else {
//			u.Data["json"] = uu
//		}
//	}
//	u.ServeJSON()
//}
