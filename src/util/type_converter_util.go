package util

import (
	"log"
	"strconv"
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
