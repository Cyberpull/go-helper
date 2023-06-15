package objects

import (
	"encoding/json"

	"cyberpull.com/gotk/validator"
)

type pJSON struct {
	validator *validator.ValidatorInstance
}

func (j *pJSON) Decode(data []byte, v any) (err error) {
	if err = json.Unmarshal(data, v); err != nil {
		return
	}

	if err = validator.Validate(v); err != nil {
		return
	}

	return
}

func (j *pJSON) Encode(v any) (value []byte, err error) {
	if err = validator.Validate(v); err != nil {
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

func NewJSON(validatorTagName ...string) *pJSON {
	return &pJSON{
		validator: validator.New(validatorTagName...),
	}
}

// ============================

var JSON pJSON

func init() {
	JSON.validator = validator.New()
}
