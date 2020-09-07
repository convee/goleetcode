package stl

import (
	"fmt"
	"strconv"
)

func main() {
	// byte转数字
	s = "12345"                         // s[0] 类型是byte
	num := int(s[0] - '0')              // 1
	str := string(s[0])                 // "1"
	b := byte(num + '0')                // '1'
	fmt.Printf("%d%s%c\n", num, str, b) // 111

	// 字符串转数字
	num, _ := strconv.Atoi()
	str := strconv.Itoa()
}
