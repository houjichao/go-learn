package myConst

import (
	"fmt"
	"unsafe"
)

/*
  常量的定义格式：const identifier [type] = value
  显式类型定义： const b string = "abc"
  隐式类型定义： const b = "abc"
  多个相同类型的声明可以简写为：
  const c_name1, c_name2 = value1, value2
*/
//常量用作枚举
const (
	Unknown = 0
	Female = 1
	Male = 2
)

const (
	//常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过
	e = "abc"
	f = len(e)
	//unsafe.Sizeof返回的是数据类型的大小，而string在Go中并不是直存类型，它是一个结构体类型：
	g = unsafe.Sizeof(e)
)
type StringHeader struct {
	//在64位系统上uintptr和int都是8字节，加起来就16了。
	Data uintptr
	Len  int
}

func ConstTest1() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d", area)
	println()
	println(a, b, c)

	println(e,f,g)
}
