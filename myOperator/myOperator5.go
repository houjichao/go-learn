package myOperator

import "fmt"

/*
赋值运算符
= 简单的赋值运算符，将一个表达式的值赋给一个左值
+= 相加后再赋值
-= 相减后再赋值
*= 相乘后再赋值
/= 相除后再赋值
%= 求余后再赋值
<<= 左移后赋值
>>= 右移后赋值
&= 按位与后赋值
^= 按位异或后赋值
|= 按位或后赋值
*/
func OperatorDemo5() {
	var a int = 21
	var c int

	c = a
	fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c)

	c += a
	fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c)

	c -= a
	fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c)

	c *= a
	fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c)

	c /= a
	fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c)

	c = 200

	c <<= 2
	fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c)

	c >>= 2
	fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c)

	c &= 2
	fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c)

	c ^= 2
	fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c)

	c |= 2
	fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c)

}
