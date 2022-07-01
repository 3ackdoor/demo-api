package null

import (
	"database/sql"
	"encoding/json"
	"strings"
	"time"
)

type String struct {
	sql.NullString
}

// NullString MarshalJSON interface redefinition
func (r String) MarshalJSON() ([]byte, error) {
	if r.Valid {
			return json.Marshal(r.String)
	} else {
			return json.Marshal(nil)
	}
}

func (r *String) UnmarshalJSON(stringBytes []byte) error {
	if len(stringBytes) == 1 && stringBytes[0] == 0 {
		r.Valid = true
		return nil
	}

	s := strings.Trim(string(stringBytes), "\"")
	if s == "null" {
		r.Valid = true
		return nil
	}

	var data string
	if err := json.Unmarshal(stringBytes, &data); err != nil {
		r.Valid = false
		return err
	}

	r.Valid = true
	r.String = data

	return nil
}

type Int64 struct {
	sql.NullInt64
}

// NullInt64 MarshalJSON interface redefinition
func (r Int64) MarshalJSON() ([]byte, error) {
	if r.Valid {
			return json.Marshal(r.Int64)
	} else {
			return json.Marshal(nil)
	}
}

func (r *Int64) UnmarshalJSON(stringBytes []byte) error {
	if len(stringBytes) == 1 && stringBytes[0] == 0 {
		r.Valid = true
		return nil
	}

	s := strings.Trim(string(stringBytes), "\"")
	if s == "null" {
		r.Valid = true
		return nil
	}

	var data int64
	if err := json.Unmarshal(stringBytes, &data); err != nil {
		r.Valid = false
		return err
	}

	r.Valid = true
	r.Int64 = data

	return nil
}

type Int32 struct {
	sql.NullInt32
}

// NullInt32 MarshalJSON interface redefinition
func (r Int32) MarshalJSON() ([]byte, error) {
	if r.Valid {
			return json.Marshal(r.Int32)
	} else {
			return json.Marshal(nil)
	}
}

func (r *Int32) UnmarshalJSON(stringBytes []byte) error {
	if len(stringBytes) == 1 && stringBytes[0] == 0 {
		r.Valid = true
		return nil
	}

	s := strings.Trim(string(stringBytes), "\"")
	if s == "null" {
		r.Valid = true
		return nil
	}

	var data int32
	if err := json.Unmarshal(stringBytes, &data); err != nil {
		r.Valid = false
		return err
	}

	r.Valid = true
	r.Int32 = data

	return nil
}


type Int16 struct {
	sql.NullInt16
}

// NullInt16 MarshalJSON interface redefinition
func (r Int16) MarshalJSON() ([]byte, error) {
	if r.Valid {
			return json.Marshal(r.Int16)
	} else {
			return json.Marshal(nil)
	}
}

func (r *Int16) UnmarshalJSON(stringBytes []byte) error {
	if len(stringBytes) == 1 && stringBytes[0] == 0 {
		r.Valid = true
		return nil
	}

	s := strings.Trim(string(stringBytes), "\"")
	if s == "null" {
		r.Valid = true
		return nil
	}

	var data int16
	if err := json.Unmarshal(stringBytes, &data); err != nil {
		r.Valid = false
		return err
	}

	r.Valid = true
	r.Int16 = data

	return nil
}

type Byte struct {
	sql.NullByte
}

// NullByte MarshalJSON interface redefinition
func (r Byte) MarshalJSON() ([]byte, error) {
	if r.Valid {
			return json.Marshal(r.Byte)
	} else {
			return json.Marshal(nil)
	}
}

func (r *Byte) UnmarshalJSON(stringBytes []byte) error {
	if len(stringBytes) == 1 && stringBytes[0] == 0 {
		r.Valid = true
		return nil
	}

	s := strings.Trim(string(stringBytes), "\"")
	if s == "null" {
		r.Valid = true
		return nil
	}

	var data byte
	if err := json.Unmarshal(stringBytes, &data); err != nil {
		r.Valid = false
		return err
	}

	r.Valid = true
	r.Byte = data

	return nil
}

type Float64 struct {
	sql.NullFloat64
}

// NullFloat64 MarshalJSON interface redefinition
func (r Float64) MarshalJSON() ([]byte, error) {
	if r.Valid {
			return json.Marshal(r.Float64)
	} else {
			return json.Marshal(nil)
	}
}

func (r *Float64) UnmarshalJSON(stringBytes []byte) error {
	if len(stringBytes) == 1 && stringBytes[0] == 0 {
		r.Valid = true
		return nil
	}

	s := strings.Trim(string(stringBytes), "\"")
	if s == "null" {
		r.Valid = true
		return nil
	}

	var data float64
	if err := json.Unmarshal(stringBytes, &data); err != nil {
		r.Valid = false
		return err
	}

	r.Valid = true
	r.Float64 = data

	return nil
}

type Bool struct {
	sql.NullBool
}

// NullBool MarshalJSON interface redefinition
func (r Bool) MarshalJSON() ([]byte, error) {
	if r.Valid {
			return json.Marshal(r.Bool)
	} else {
			return json.Marshal(nil)
	}
}

func (r *Bool) UnmarshalJSON(stringBytes []byte) error {
	if len(stringBytes) == 1 && stringBytes[0] == 0 {
		r.Valid = true
		return nil
	}

	s := strings.Trim(string(stringBytes), "\"")
	if s == "null" {
		r.Valid = true
		return nil
	}

	var data bool
	if err := json.Unmarshal(stringBytes, &data); err != nil {
		r.Valid = false
		return err
	}

	r.Valid = true
	r.Bool = data

	return nil
}

type Time struct {
	sql.NullTime
}

// NullTime MarshalJSON interface redefinition
func (r Time) MarshalJSON() ([]byte, error) {
	if r.Valid {
			return json.Marshal(r.Time)
	} else {
			return json.Marshal(nil)
	}
}

func (r *Time) UnmarshalJSON(stringBytes []byte) error {
	if len(stringBytes) == 1 && stringBytes[0] == 0 {
		r.Valid = true
		return nil
	}

	s := strings.Trim(string(stringBytes), "\"")
	if s == "null" {
		r.Valid = true
		return nil
	}

	var data time.Time
	if err := json.Unmarshal(stringBytes, &data); err != nil {
		r.Valid = false
		return err
	}

	r.Valid = true
	r.Time = data

	return nil
}

