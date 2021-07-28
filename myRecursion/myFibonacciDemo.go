package myRecursion

import "fmt"

/**
递归，就是在运行的过程中调用自己。
Go 语言支持递归。但我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中。
递归函数对于解决数学上的问题是非常有用的，就像计算阶乘，生成斐波那契数列等。

斐波那契数列（Fibonacci sequence），又称黄金分割数列，因数学家莱昂纳多·斐波那契（Leonardoda Fibonacci）以兔子繁殖为例子而引入，
故又称为“兔子数列”，指的是这样一个数列：0、1、1、2、3、5、8、13、21、34、
这个数列从第3项开始，每一项都等于前两项之和。
*/
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-2) + Fibonacci(n-1)
}

func CallFibonacci() {
	fmt.Println("----------fibonacci demo start----------")
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", Fibonacci(i))
	}
	fmt.Println()
	fmt.Println("----------fibonacci demo end----------")
}
