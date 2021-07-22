package myDefer

import "fmt"

/*
Go语言的 defer 语句会将其后面跟随的语句进行延迟处理，在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 的逆序进行执行，
也就是说，先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行。

关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源，典型的例子就是对一个互斥解锁，或者关闭一个文件。
*/
func DeferDemo() {

	fmt.Println("defer begin")
	defer fmt.Println("----------defer demo end----------")
	// 将defer放入延迟调用栈
	defer fmt.Println(1)
	defer fmt.Println(2)
	// 最后一个放入, 位于栈顶, 最先调用
	defer fmt.Println(3)
	fmt.Println("defer end")

	defer fmt.Println("----------defer demo start----------")

	/*
		结果分析如下：
		代码的延迟顺序与最终的执行顺序是反向的。
		延迟调用是在 defer 所在函数结束时进行，函数结束可以是正常返回时，也可以是发生宕机时。
	*/

}
