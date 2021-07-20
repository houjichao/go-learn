package myOperator

import "fmt"

/*
其他运算符
& 返回变量存储地址  &a; 将给出变量的实际地址。
* 指针变量  *a; 是一个指针变量

指针变量 * 和地址值 & 的区别：指针变量保存的是一个地址值，会分配独立的内存来存储一个整型数字。
当变量前面有 * 标识时，才等同于 & 的用法，否则会直接输出一个整型数字。
*/
func OperatorDemo6() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b)
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c)

	/*  & 和 * 运算符实例 */
	ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a)
	fmt.Println("ptr:", ptr) //0xc0000b4038
	fmt.Printf("*ptr 为 %d\n", *ptr)

	//使用指针变量与不使用的区别：
	var f int = 4
	var ptr1 int
	ptr1 = f
	fmt.Println("ptr1:", ptr1) //4
	f = 15
	fmt.Println("ptr1:", ptr1) //4

	var g int = 5
	var ptr2 *int
	ptr2 = &g
	fmt.Println("ptr2:", *ptr2) //5
	g = 15
	fmt.Println("ptr2:", *ptr2) //15

}
