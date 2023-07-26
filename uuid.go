package gotk

import (
	"cyberpull.com/gotk/v3/errors"

	"github.com/google/uuid"
)

func UUID() (value string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.From(r)
		}
	}()

	value = uuid.NewString()

	return
}
