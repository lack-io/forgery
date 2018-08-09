package forgery

import (
	"math/rand"
	"time"
)

var (
	hex string = "0123456789ABCDEF"
	lowercase string = "abcdefghijklmnopqrstuvwxyz"
	uppercase string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits string = "0123456"
)


func HexColor() (string, error) {
	result := ""
	for i := 0; i < 6; i++ {
		rand.Seed(time.Now().UnixNano())
		length := len(hex)
		rd := rand.Intn(length)
		result += hex[rd:rd+1]
	}
	return result, nil
}

func HexColorShort() (string, error) {
	result := ""
	for i := 0; i < 3; i++ {
		rand.Seed(time.Now().UnixNano())
		length := len(hex)
		rd := rand.Intn(length)
		result += hex[rd:rd+1]
	}
	return result, nil
}

func Text() (string, error) {
	result := ""
	basicText := lowercase + uppercase + " " + digits
	rand.Seed(time.Now().UnixNano())
	length := len(basicText)
	rd := rand.Intn(5) + 15
	for i := 0; i < rd; i++ {
		rand.Seed(time.Now().UnixNano())
		rdIndex := rand.Intn(length)
		result += basicText[rdIndex:rdIndex+1]
	}
	return result, nil
}