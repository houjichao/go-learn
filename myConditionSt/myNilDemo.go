package myConditionSt

import "fmt"

/*
//nil是一个预先声明的标识符，表示指针、通道、函数、接口、映射或切片类型。
在Go语言中，如果你声明了一个变量但是没有对它进行赋值操作，那么这个变量就会有一个类型的默认零值。这是每种类型对应的零值
nil是go语言中预先的标识符
我们可以直接使用nil，而不用声明它。

nil可以代表很多类型的零值
在go语言中，nil可以代表下面这些类型的零值：

指针类型（包括unsafe中的）
map类型
slice类型
function类型
channel类型
interface类型

*/
func NilDemo() {
	var a []int
	if a == nil {
		fmt.Println(123)
	}
	fmt.Println(a)

	b := []int{}
	if b == nil {
		fmt.Println("b为nil")
	} else {
		fmt.Println("b不是nil")
	}
	fmt.Println(b)
	/*
		区别两者之间的关系：
		 在第一个代码块中，a声明了变量但没有进行赋值，因此为nil
		 在第二个代码块中，a声明了变量并进行了初始化操作，只不过里面是空的，因此就行了赋值，不为nil
	*/

}
