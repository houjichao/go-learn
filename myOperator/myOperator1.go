package myOperator

import "fmt"

/*
算术运算符

+ - * /
%  求余
++ 自增
-- 自减
*/
func OperatorDemo1() {
	fmt.Println("算术运算符")

	var a int = 21
	var b int = 20
	var c int
	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)

	a++ // 这是允许的，类似 a = a + 1,结果与 a++ 相同
	a-- //与 a++ 相似
	//a = a++ // 这是不允许的，会出现编译错误 syntax error: unexpected ++ at end of statement

}
