package validate

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/stikkas/integrator/internal/handler/model"
	"github.com/stikkas/integrator/internal/utils"
	"strings"
)

func Study(err error) []*model.FieldError {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		errs := make([]*model.FieldError, len(ve))
		for i, fe := range ve {
			parts := strings.Split(fe.Namespace(), ".")
			var item string
			if size := len(parts); size > 2 {
				item = parts[len(parts)-2]
			} else {
				item = parts[0]
			}
			val := validators[fe.Tag()]
			var message string
			if val != nil {
				message = val.Message
			} else {
				message = "неизвестная ошибка"
			}
			errs[i] = &model.FieldError{
				Item:          utils.FirstToLower(item),
				Field:         utils.FirstToLower(fe.Field()),
				Message:       message,
				RejectedValue: fe.Value(),
			}
		}
		return errs
	}
	return nil
}
