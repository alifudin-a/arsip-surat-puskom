package helper

import (
	"database/sql/driver"
	"errors"
)

type NullString string

func (s *NullString) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}
	strVal, ok := value.(string)
	if !ok {
		return errors.New("column is not a string")
	}
	*s = NullString(strVal)
	return nil
}

func (s NullString) Value() (driver.Value, error) {
	if len(s) == 0 { // if nil or empty string
		return nil, nil
	}
	return string(s), nil
}

type NullInt64 int64

func (i *NullInt64) Scan(value interface{}) error {
	if value == nil {
		*i = 0
		return nil
	}

	intVal, ok := value.(int64)
	if !ok {
		return errors.New("column is not a integer")
	}

	*i = NullInt64(intVal)
	return nil
}

func (i NullInt64) Value() (driver.Value, error) {
	if i == 0 {
		return nil, nil
	}

	return int64(i), nil
}
