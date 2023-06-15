package validator

var instance *ValidatorInstance

func Validate(data any, rules ...any) (err error) {
	return instance.Validate(data, rules...)
}

func SetTagName(name string) {
	instance.validate.SetTagName(name)
}

func ResetTagName() {
	instance.validate.SetTagName(defaultTagName)
}

// ==========================

func one[T any](def T, attr []any) T {
	if len(attr) == 0 {
		return attr[0].(T)
	}

	return def
}

// ==========================

func init() {
	instance = New()
}
