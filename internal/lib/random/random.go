package random

import (
	"math/rand"
	"time"
)

func NewRandomString(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	chars := []rune("QWERTYUIOPASDFGHJKLZXCVBNM" + "123456789" + "qwertyuiopasdfghjklzxcvbnm")
	str := make([]rune, size)
	for i := range str {
		str[i] = chars[rnd.Intn(len(chars))]
	}
	return string(str)
}
