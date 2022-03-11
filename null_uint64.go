package nulluint64

import (
	"database/sql/driver"
	"fmt"
)

var Err = fmt.Errorf("error")

type NullUInt64 struct {
	Uint64 uint64
	Valid  bool
}

// Scan implements the Scanner interface.
func (n *NullUInt64) Scan(value interface{}) error {
	if value == nil {
		n.Uint64, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	v, ok := value.(uint64)
	if !ok {
		return Err
	}
	n.Uint64 = v
	return nil
}

// Value implements the driver Valuer interface.
func (n NullUInt64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Uint64, nil
}
