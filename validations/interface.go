package validations

import (
	"ApiGuide/models"
	"github.com/astaxie/beego/validation"
)

func CheckInterface(inter *models.Interface) []interface{} {
	var errorList []interface{}
	valid := validation.Validation{}
	valid.Required(inter.Name, "name")
	valid.Required(inter.Path, "path")
	valid.Required(inter.Version, "version")
	valid.Required(inter.Require, "require")
	valid.Required(inter.Example, "example")
	valid.Required(inter.Type, "type")
	valid.Required(inter.Desc, "desc")
	valid.Required(inter.Modular, "modular")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			errorList = append(errorList, err)
		}
	}
	return errorList
}
