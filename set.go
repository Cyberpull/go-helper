package gotk

import (
	"fmt"
	"strconv"
	"strings"

	"cyberpull.com/gotk/v2/errors"

	"golang.org/x/exp/constraints"
)

type SetConstraint interface {
	constraints.Integer | constraints.Float | ~string
}

func Join[T SetConstraint](entries []T, delim string) (value string, err error) {
	for _, entry := range entries {
		data := fmt.Sprint(entry)
		value += delim + data
	}

	value = strings.TrimPrefix(value, delim)

	return
}

func Split[T SetConstraint](data string, delim string) (value []T, err error) {
	var t T

	value = make([]T, 0)

	for _, entry := range strings.Split(data, delim) {
		var newValue any

		switch any(t).(type) {
		case string:
			newValue = entry

		case int, int8, int16, int32, int64:
			newValue, err = strconv.ParseInt(entry, 0, 64)

		case uint, uint8, uint16, uint32, uint64:
			newValue, err = strconv.ParseUint(entry, 0, 64)

		case float32, float64:
			newValue, err = strconv.ParseFloat(entry, 64)

		default:
			err = errors.Newf("Type '%T' not allowed", t)
		}

		if err != nil {
			return
		}

		value = append(value, any(newValue).(T))
	}

	return
}
