package validations

import (
	"ApiGuide/models"
	"github.com/astaxie/beego/validation"
)

func CheckModular(modular *models.Modular) []interface{} {
	var errorList []interface{}
	valid := validation.Validation{}
	valid.Required(modular.Name, "name")
	valid.Required(modular.Project, "project")
	valid.Required(modular.Prefix, "prefix")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			errorList = append(errorList, err)
		}
	}
	return errorList
}
