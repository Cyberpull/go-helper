package validator

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

type ValidatorInstance struct {
	validate *validator.Validate
}

func (i *ValidatorInstance) Validate(data any, rules ...any) (err error) {
	rValue := reflect.ValueOf(data)

	rType := rValue.Type()
	rKind := rType.Kind()

	if rKind == reflect.Pointer {
		rValue = rValue.Elem()
	}

	value := rValue.Interface()

	switch rKind {
	case reflect.Struct:
		return i.validate.Struct(value)
	default:
		rule := one[string]("", rules)
		return i.validate.Var(value, rule)
	}
}

// ==========================

const defaultTagName string = "validate"

func New(tagName ...string) *ValidatorInstance {
	if len(tagName) == 0 {
		tagName = append(tagName, defaultTagName)
	}

	validate := validator.New()
	validate.SetTagName(tagName[0])

	return &ValidatorInstance{
		validate: validate,
	}
}
