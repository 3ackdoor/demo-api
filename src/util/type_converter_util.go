package util

import (
	"database/sql"
	"log"
	"strconv"
	"time"
)

func ConvertStringToUint(s string) uint {
	var i uint
	if val, err := strconv.ParseUint(s, 10, 32); err != nil {
		log.Panic(err)
	} else {
		i = uint(val)
	}
	return i
}

func NewNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{
			Valid: false,
		}
	}

	if len(*s) == 0 {
		return sql.NullString{}
	}

	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}


func NewNullTime(s *time.Time) sql.NullTime {
	if s == nil {
		return sql.NullTime{
			Valid: false,
		}
	}

	return sql.NullTime{
		Time: *s,
		Valid:  true,
	}
}
