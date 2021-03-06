package myPointer

import "fmt"

/**

Go 语言中指针是很容易学习的，Go 语言中使用指针可以更简单的执行一些任务。
接下来让我们来一步步学习 Go 语言指针。
我们都知道，变量是一种使用方便的占位符，用于引用计算机内存地址。
Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。

*/
func PointerDemo1() {

	fmt.Println("----------pointer demo1 start----------")

	var a int = 10

	fmt.Printf("变量的地址：%x\n", &a)

	fmt.Println("----------pointer demo1 end----------")

}
