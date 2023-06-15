package validator

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Validate(data any, rules ...any) (err error) {
	rValue := reflect.ValueOf(data)

	rType := rValue.Type()
	rKind := rType.Kind()

	if rKind == reflect.Pointer {
		rValue = rValue.Elem()
	}

	value := rValue.Interface()

	switch rKind {
	case reflect.Struct:
		return validate.Struct(value)
	default:
		rule := one[string]("", rules)
		return validate.Var(value, rule)
	}
}

func one[T any](def T, attr []any) T {
	if len(attr) == 0 {
		return attr[0].(T)
	}

	return def
}

// ==========================

func init() {
	validate = validator.New()
}
