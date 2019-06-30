package validations

import (
	"ApiGuide/models"
	"github.com/astaxie/beego/validation"
)

func ProjectAddCheck(pro *models.Project) []interface{} {
	var errorList []interface{}
	valid := validation.Validation{}
	valid.Required(pro.Name, "name")
	valid.Required(pro.Host, "host")
	valid.Required(pro.User, "user")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			errorList = append(errorList, err)
		}
	}
	return errorList
}
