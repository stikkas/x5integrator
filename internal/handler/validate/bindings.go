package validate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/stikkas/integrator/internal/handler/model"
	"log"
	"slices"
)

func Bind() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// StudyRequest.Data[]
		err := v.RegisterValidation("data", func(fl validator.FieldLevel) bool {
			data, ok := fl.Field().Interface().([]*model.RequestOperation)
			if !ok {
				log.Printf("value is not []*model.RequrestOperation, but %v", fl.Field())
				return false
			}
			if data == nil || len(data) == 0 {
				return false
			}
			return true
		})
		if err != nil {
			log.Printf("can't register 'data' binding: %v", err)
		}

		// StudyRequest.SystemId
		err = v.RegisterValidation("systemId", func(fl validator.FieldLevel) bool {
			value := fl.Field().String()
			return slices.Contains(model.SystemIds, value)
		})
		if err != nil {
			log.Printf("can't register 'systemId' binding: %v", err)
		}

		// RequestOperation.OperationType
		err = v.RegisterValidation("operation", func(fl validator.FieldLevel) bool {
			value := model.OperationType(fl.Field().Int())
			return slices.Contains(model.OperationTypes, value)
		})
		if err != nil {
			log.Printf("can't register 'operation' binding: %v", err)
		}

		// RequestOperation.StudyType
		err = v.RegisterValidation("study", func(fl validator.FieldLevel) bool {
			value := model.StudyType(fl.Field().Int())
			return slices.Contains(model.StudyTypes, value)
		})
		if err != nil {
			log.Printf("can't register 'study' binding: %v", err)
		}

	}
}
