package main

import (
	"fmt"
	"strings"
)

//1108. IP 地址无效化
//https://leetcode-cn.com/problems/defanging-an-ip-address/
func defangIPaddr(address string) string {
	var s strings.Builder
	for i := range address {
		if address[i] == '.' {
			s.WriteString("[.]")
		} else {
			s.WriteByte(address[i])
		}
	}
	return s.String()
}

func main() {
	ip := "127.0.0.1"
	fmt.Println(defangIPaddr(ip))
}
