package validations

import (
	"ApiGuide/models"
	"github.com/astaxie/beego/validation"
)

// 注册参数验证
func RegisterCheck(u *models.User) []interface{} {
	var errorList []interface{}
	valid := validation.Validation{}
	valid.Required(u.Username, "username")
	valid.Required(u.Password, "password")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			errorList = append(errorList, err)
		}
	}
	return errorList
}

// 登陆参数验证
func LoginCheck(u *models.User) []interface{} {
	var errorList []interface{}
	valid := validation.Validation{}
	valid.Required(u.Username, "username")
	valid.Required(u.Password, "password")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			errorList = append(errorList, err)
		}
	}
	return errorList
}

// 验证userId
func UserIdCheck(user *models.User) []interface{} {
	var errorList []interface{}
	vaild := validation.Validation{}
	vaild.Required(user.Id, "id")
	if vaild.HasErrors() {
		for _, err := range vaild.Errors {
			errorList = append(errorList, err)
		}
	}
	return errorList
}
