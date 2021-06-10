package myStr

import (
	"fmt"
	"strings"
)

func Test() {
	str := "这里是   www\n.houjichao\n.com"
	fmt.Println("--------原字符串---------")
	fmt.Println(str)
	//去除空格
	str = strings.Replace(str, " ", "", -1)
	//去除换行符
	str = strings.Replace(str, "\n", "", -1)

	fmt.Println("--------去除空格与换行后----------")
	fmt.Println(str)

}
