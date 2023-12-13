package dbo

import (
	"database/sql/driver"
	"encoding/json"

	"cyberpull.com/gotk/v2"
	"cyberpull.com/gotk/v2/errors"
)

type Set[T gotk.SetConstraint] struct {
	Data []T
}

func (n *Set[T]) Scan(value any) (err error) {
	switch data := value.(type) {
	case []T:
		n.Data = data

	case string:
		n.Data, err = gotk.Split[T](data, ",")

	case []byte:
		n.Data, err = gotk.Split[T](string(data), ",")

	default:
		err = errors.New("Invalid data type")
	}

	return
}

func (n Set[T]) Value() (value driver.Value, err error) {
	value, err = gotk.Join[T](n.Data, ",")
	return
}

func (n *Set[T]) UnmarshalJSON(b []byte) error {
	value := make([]T, 0)

	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	return n.Scan(value)
}

func (n Set[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Data)
}
