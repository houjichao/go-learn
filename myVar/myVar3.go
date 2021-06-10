package myVar

import "fmt"

/*
  空白标识符在函数返回值时的使用
*/
func VarTest3() {
	_,numb,strs := numbers() //只获取函数返回值的后两个
	fmt.Println(numb,strs)

	numa, numb, _ := numbers()
	fmt.Println(numa, numb)

}

//一个可以返回多个值的函数
func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}
