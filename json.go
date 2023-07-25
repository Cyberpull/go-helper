package gotk

import (
	"encoding/json"
)

type JSONEngine interface {
	Decode(data []byte, v any) (err error)
	Encode(v any) (value []byte, err error)
}

// ============================

type jsonEngine struct {
	validator Validator
}

func (e *jsonEngine) Decode(data []byte, v any) (err error) {
	if err = json.Unmarshal(data, v); err != nil {
		return
	}

	if err = e.validator.Validate(v); err != nil {
		return
	}

	return
}

func (e *jsonEngine) Encode(v any) (value []byte, err error) {
	if err = e.validator.Validate(v); err != nil {
		return
	}

	value, err = json.Marshal(v)

	return
}

// ============================

func ParseJSON(data []byte, v any) (err error) {
	return JSON.Decode(data, v)
}

func ToJSON(v any) (value []byte, err error) {
	return JSON.Encode(v)
}

// ============================

func NewJSON(validatorTagName ...string) JSONEngine {
	engine := new(jsonEngine)
	prepareJSON(engine, validatorTagName...)
	return engine
}

// ============================

var JSON jsonEngine

func prepareJSON(engine *jsonEngine, validatorTagName ...string) {
	engine.validator = NewValidator(validatorTagName...)
}

func init() {
	prepareJSON(&JSON)
}
