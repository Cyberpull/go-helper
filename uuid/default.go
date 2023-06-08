package uuid

import (
	"cyberpull.com/gotk/errors"

	"github.com/google/uuid"
)

func Generate() (value string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.From(r)
		}
	}()

	value = uuid.NewString()

	return
}
