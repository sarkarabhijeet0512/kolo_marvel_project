package utils

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"strconv"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
func ConvertStringToInt(numeric string) int {
	convertNumber, err := strconv.Atoi(numeric)
	if err != nil {
		log.Println("Error converting string to integer: ", err)
		return 0
	} else {
		return convertNumber
	}
}
