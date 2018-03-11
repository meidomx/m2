package edcode

import (
	"strings"
	"strconv"
)

func Trim(src string) string {
	return strings.TrimSpace(src)
}

//ReverseStr
func ReverseStrBytes(src string) string {
	b:= []byte(src)
	i := 0
	l := len(b)
	nb := make([]byte, l)
	copy(nb, b)
	l -= 1
	for true {
		if i >= l {
			break
		}
		nb[i], nb[l] = nb[l], nb[i]
		i++
		l--
	}
	return string(nb)
}

func IntToStr(i uint32) string {
	return strconv.Itoa(int(i))
}
