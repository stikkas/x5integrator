package validate

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/stikkas/integrator/internal/handler/model"
	"github.com/stikkas/integrator/internal/utils"
	"strings"
)

func Study(err error, requestUri string) *model.ProblemDetail {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		detail := model.ProblemDetail{
			Type:     "https://wiki.x5.ru/pages/viewpage.action?pageId=1302077850",
			Instance: requestUri,
			Title:    "Входные параметры не соответствуют требованиям",
			Errors:   make([]*model.FieldError, len(ve)),
		}
		for i, fe := range ve {
			parts := strings.Split(fe.Namespace(), ".")
			var item string
			if size := len(parts); size > 2 {
				item = parts[len(parts)-2]
			} else {
				item = parts[0]
			}
			detail.Errors[i] = &model.FieldError{
				Item:          utils.FirstToLower(item),
				Field:         utils.FirstToLower(fe.Field()),
				Message:       getErrorMsg(fe),
				RejectedValue: fe.Value(),
			}
		}
		return &detail
	}
	return nil
}

func getErrorMsg(ve validator.FieldError) string {
	switch ve.Tag() {
	case "required":
		switch ve.Field() {
		case "Tn":
			return "не задан ТН сотрудника"
		case "StudyId":
			return "не задан идентификатор учебной сущности"
		case "RequestUUID":
			return "не задан UUID запроса"
		}
	case "systemId":
		return "не задан или недопустимое значение идентификатора системы-инициатора"
	case "data":
		return "не заданы операции"
	case "operation":
		return "не задан или неподдерживаемый тип операции"
	case "study":
		return "не задан или неподдерживаемый тип учебной сущности"
	}
	return "неизвестная ошибка"

	//RequestOperationRaw::tn.name,
	//String?::isLessThanOne to "ТН сотрудника меньше 1",

	//	RequestOperationRaw::studyId.name,
	//	String?::isLessThanOne to "идентификатор учебной сущности меньше 1",
	//	String?::isNotDate to "формат даты не соответствует паттерну $DATE_PATTERN"
	//if (target.operationType?.toIntOrNull() == 2) {
	//	dateChecks.add(0, String?::isNullOrBlank to "не задана дата назначения обучения")
	//}
	//	RequestOperationRaw::date.name,
	//	*dateChecks.toTypedArray()
}
