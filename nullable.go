package nullable

import (
	"database/sql"
	"encoding/json"
)

// Bool represents a bool that may be null.
// Similar to sql.NullBool, but implement json.Marshaler/json.Unmarshaler
type Bool struct {
	sql.NullBool
}

// MarshalJSON implements the json.Marshaler interface.
func (b Bool) MarshalJSON() ([]byte, error) {
	return marshalJSON(b.Valid, b.Bool)
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (b *Bool) UnmarshalJSON(data []byte) error {
	var x *bool
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		b.Valid = true
		b.Bool = *x
	} else {
		b.Valid = false
	}
	return nil
}

// Float64 represents a float64 that may be null.
// Similar to sql.NullFloat64, but implement json.Marshaler/json.Unmarshaler
type Float64 struct {
	sql.NullFloat64
}

// MarshalJSON implements the json.Marshaler interface.
func (f Float64) MarshalJSON() ([]byte, error) {
	if f.Valid {
		return json.Marshal(f.Float64)
	} else {
		return json.Marshal(nil)
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (f *Float64) UnmarshalJSON(data []byte) error {
	var x *float64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		f.Valid = true
		f.Float64 = *x
	} else {
		f.Valid = false
	}
	return nil
}

// Int64 represents an int64 that may be null.
// Similar to sql.NullInt64, but implement json.Marshaler/json.Unmarshaler
type Int64 struct {
	sql.NullInt64
}

// MarshalJSON implements the json.Marshaler interface.
func (v Int64) MarshalJSON() ([]byte, error) {
	return marshalJSON(v.Valid, v.Int64)
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (v *Int64) UnmarshalJSON(data []byte) error {
	var x *int64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int64 = *x
	} else {
		v.Valid = false
	}
	return nil
}

// String represents a string that may be null.
// Similar to sql.NullString, but implement json.Marshaler/json.Unmarshaler
type String struct {
	sql.NullString
}

// MarshalJSON implements the json.Marshaler interface.
func (s String) MarshalJSON() ([]byte, error) {
	return marshalJSON(s.Valid, s.String)
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (s *String) UnmarshalJSON(data []byte) error {
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		s.Valid = true
		s.String = *x
	} else {
		s.Valid = false
	}
	return nil
}

func marshalJSON(valid bool, v interface{}) ([]byte, error) {
	if valid {
		return json.Marshal(v)
	}
	return json.Marshal(nil)
}
