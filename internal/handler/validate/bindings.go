package validate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/stikkas/integrator/internal/handler/model"
	"log"
	"reflect"
	"slices"
	"strings"
)

type safeGuard struct {
	Tag      string
	Validate func(fl validator.FieldLevel) bool
	Message  string
}

const datePattern = "yyyy-MM-dd"

var dataGuard = safeGuard{
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

var systemIdGuard = safeGuard{
	"systemId",
	func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		return slices.Contains(model.SystemIds, value)
	},
	"не задан или недопустимое значение идентификатора системы-инициатора",
}

var operationGuard = safeGuard{
	"operation",
	func(fl validator.FieldLevel) bool {
		value := model.OperationType(fl.Field().Int())
		return slices.Contains(model.OperationTypes, value)
	},
	"не задан или неподдерживаемый тип операции",
}

var studyGuard = safeGuard{
	"study",
	func(fl validator.FieldLevel) bool {
		value := model.StudyType(fl.Field().Int())
		return slices.Contains(model.StudyTypes, value)
	},
	"не задан или неподдерживаемый тип учебной сущности",
}

var tnGuard = safeGuard{
	"tn",
	func(fl validator.FieldLevel) bool {
		return fl.Field().Uint() > 0
	},
	"не задан или неправильный ТН сотрудника",
}

var uuidGuard = safeGuard{
	"ruuid",
	func(fl validator.FieldLevel) bool {
		return len(strings.Trim(fl.Field().String(), " ")) > 0
	},
	"не задан UUID запроса",
}

var studyIdGuard = safeGuard{
	"studyId",
	func(fl validator.FieldLevel) bool {
		return fl.Field().Uint() > 0
	},
	"не задан или неправильный идентификатор учебной сущности",
}

var dateGuard = safeGuard{
	"date",
	func(fl validator.FieldLevel) bool {
		fs := reflect.Indirect(fl.Parent()).FieldByName("OperationType")
		if int(fs.Int()) == int(model.Study) && fl.Field().IsZero() {
			return false
		}
		return true
	},
	"не задана дата назначения обучения или формат не соответствует паттерну " + datePattern,
}

var validators = map[string]*safeGuard{
	dataGuard.Tag:      &dataGuard,
	systemIdGuard.Tag:  &systemIdGuard,
	operationGuard.Tag: &operationGuard,
	studyGuard.Tag:     &studyGuard,
	tnGuard.Tag:        &tnGuard,
	studyIdGuard.Tag:   &studyIdGuard,
	uuidGuard.Tag:      &uuidGuard,
	dateGuard.Tag:      &dateGuard,
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
