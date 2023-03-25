package numUtils

import (
	"log"
	"strconv"
)

func ParseUint(str string) uint {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		log.Println("Error:", err)
		return 0
	}
	return uint(num)
}
