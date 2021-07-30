package myArray

import "fmt"

/*
Go 语言提供了数组类型的数据结构。
数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整型、字符串或者自定义类型。
相对于去声明 number0, number1, ..., number99 的变量，使用数组形式 numbers[0], numbers[1] ..., numbers[99] 更加方便且易于扩展。
数组元素可以通过索引（位置）来读取（或者修改），索引从 0 开始，第一个元素索引为 0，第二个索引为 1，以此类推。

声明数组
Go 语言数组声明需要指定元素类型及元素个数，语法格式如下：
var variable_name [SIZE] variable_type
以上为一维数组的定义方式。例如以下定义了数组 balance 长度为 10 类型为 float32：
var balance [10] float32

*/
func ArrayDemo1() {

	fmt.Println("----------array demo1 start----------")

	//声明数组
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	//长度不确定可以
	var balance1 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	var i int
	for i = 0; i < 5; i++ {
		fmt.Printf("balance element [%d] = %f \n", i, balance[i])
	}

	for i = 0; i < 5; i++ {
		fmt.Printf("balance1 element [%d] = %f \n", i, balance1[i])
	}

	//定义长度为10的数组
	var n [10]int
	var j int
	//为数组n初始化元素
	for j = 0; j < 10; j++ {
		n[j] = j + 100
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("n Element[%d] = %d\n", j, n[j] )
	}

	fmt.Println("----------array demo1 end----------")
}
