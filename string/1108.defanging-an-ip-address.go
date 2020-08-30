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

func defangIPaddr2(address string) string {
	var res []byte
	for _, v := range address {
		if v == '.' {
			res = append(res, '[', '.', ']')
		} else {
			res = append(res, byte(v))
		}
	}
	return string(res)
}

func defangIPaddr3(address string) string {
	var s []string
	for _, i := range address {
		if i == '.' {
			s = append(s, string("[.]"))
		} else {
			s = append(s, string(i))
		}
	}
	return strings.Join(s, "")
}

func main() {
	ip := "127.0.0.1"
	fmt.Println(defangIPaddr(ip))
}
