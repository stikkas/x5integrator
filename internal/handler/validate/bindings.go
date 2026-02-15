package validate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/stikkas/integrator/internal/handler/model"
	"log"
	"slices"
)

type Validator struct {
	Tag      string
	Validate func(fl validator.FieldLevel) bool
	Message  string
}

var DataValidator = Validator{
	"data",
	func(fl validator.FieldLevel) bool {
		data, ok := fl.Field().Interface().([]*model.RequestOperation)
		if !ok {
			log.Printf("value is not []*model.RequrestOperation, but %v", fl.Field())
			return false
		}
		if data == nil || len(data) == 0 {
			return false
		}
		return true
	},
	"не заданы операции",
}

var SystemIdValidator = Validator{
	"systemId",
	func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		return slices.Contains(model.SystemIds, value)
	},
	"не задан или недопустимое значение идентификатора системы-инициатора",
}

var OperationValidator = Validator{
	"operation",
	func(fl validator.FieldLevel) bool {
		value := model.OperationType(fl.Field().Int())
		return slices.Contains(model.OperationTypes, value)
	},
	"не задан или неподдерживаемый тип операции",
}

var StudyValidator = Validator{
	"study",
	func(fl validator.FieldLevel) bool {
		value := model.StudyType(fl.Field().Int())
		return slices.Contains(model.StudyTypes, value)
	},
	"не задан или неподдерживаемый тип учебной сущности",
}
var validators = map[string]*Validator{
	DataValidator.Tag:      &DataValidator,
	SystemIdValidator.Tag:  &SystemIdValidator,
	OperationValidator.Tag: &OperationValidator,
	StudyValidator.Tag:     &StudyValidator,
}

func Bind() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for _, val := range validators {
			err := v.RegisterValidation(val.Tag, val.Validate)
			if err != nil {
				log.Printf("can't register '%s' binding: %v", val.Tag, err)
			}
		}
	}
}
