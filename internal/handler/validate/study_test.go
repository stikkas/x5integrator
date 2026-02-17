package validate

import (
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
	"testing"
)

type MockError struct {
	tag       string
	namespace string
	field     string
	value     interface{}
}

func TestStudy(t *testing.T) {
	errors := []validator.FieldError{
		MockError{
			"systemId",
			"studyRequest",
			"SystemId",
			"",
		},
	}
	res := Study(validator.ValidationErrors(errors))
	expected := len(errors)
	if len(res) != expected {
		t.Errorf("wrong amount of errors: %d instead of %d", len(res), expected)
	}

	for _, err := range res {
		if err.Field != "systemId" {
			t.Errorf("expected: systemId, but got %s", err.Field)
		}
		if err.Message != systemIdGuard.Message {
			t.Errorf("expected: %s, but got %s", systemIdGuard.Message, err.Message)
		}
		if err.Item != "studyRequest" {
			t.Errorf("expected: studyRequest, but got %s", err.Item)
		}
		if err.RejectedValue != "" {
			t.Errorf("expected: , but got %s", err.RejectedValue)
		}
	}
}

func (me MockError) ActualTag() string {
	//TODO implement me
	panic("implement me")
}

func (me MockError) Namespace() string {
	return me.namespace
}

func (me MockError) StructNamespace() string {
	//TODO implement me
	panic("implement me")
}

func (me MockError) Field() string {
	return me.field
}

func (me MockError) StructField() string {
	//TODO implement me
	panic("implement me")
}

func (me MockError) Value() interface{} {
	return me.value
}

func (me MockError) Param() string {
	//TODO implement me
	panic("implement me")
}

func (me MockError) Kind() reflect.Kind {
	//TODO implement me
	panic("implement me")
}

func (me MockError) Type() reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (me MockError) Translate(ut ut.Translator) string {
	//TODO implement me
	panic("implement me")
}

func (me MockError) Error() string {
	//TODO implement me
	panic("implement me")
}

func (me MockError) Tag() string {
	return me.tag
}
