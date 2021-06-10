package myVar

import (
	"fmt"
	"strconv"
)

func VarTest2() {
	//变量声明
	//第一种：指定变量类型，如果没有初始化，则变量默认为零值。
	var i int
	var f float64
	var b bool
	var s string
	//占位符说明
	//%v 相应值的默认格式
	//%q 双引号围绕的字符串，由Go语法安全地转义
	//%d 十进制表示
	fmt.Printf("%d\n", 0x12)
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	//第二种，根据值自行判定变量类型。
	var d = true
	fmt.Println("第二种，根据值自行判定变量类型:" + strconv.FormatBool(d))

	//第三种，省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误，格式：v_name := value

	//var intVal int
	//intVal := 1 //这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明
	intVal := 1
	fmt.Println("第三种，省略 var:", intVal)

}
