package myConst

import (
	"fmt"
	"strconv"
	"unsafe"
)

const (
	a = "hello"
	m = unsafe.Sizeof(a)
)

func ConstTest4() {
	a := "hello"
	size := unsafe.Sizeof(a)
	fmt.Println("a size:" + strconv.Itoa(int(size)))
	fmt.Println("a size:", m)

}
