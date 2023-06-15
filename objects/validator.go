package objects

import "cyberpull.com/gotk/validator"

type pValidator struct {
	//
}

func (v *pValidator) SetTagName(name string) {
	validator.SetTagName(name)
}

func (v *pValidator) ResetTagName() {
	validator.ResetTagName()
}

// ========================

var Validator pValidator
