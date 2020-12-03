package utils

import (
	"log"
	"strconv"
)

//StrConvertUint convert string to uint
func StrConvertUint(str string) uint {
	u64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		log.Println(err)
		return 0
	}
	wd := uint(u64)
	return wd
}
