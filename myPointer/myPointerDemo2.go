package myPointer

import "fmt"

/**

什么是指针
一个指针变量指向了一个值的内存地址。
类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：

var var_name *var-type
var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针。以下是有效的指针声明：

var ip *int         //指向整型
var fp *float32     //指向浮点型
本例中这是一个指向 int 和 float32 的指针。

如何使用指针
指针使用流程：
1.定义指针变量。
2.为指针变量赋值。
3.访问指针变量中指向地址的值。

在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。

*/
func PointerDemo2() {
	fmt.Println("----------pointer demo2 start----------")

	var a int = 20 //声明实际变量
	var ip *int    //声明指针变量
	ip = &a        //指针变量的存储地址
	fmt.Printf("a 变量的地址是：%x\n", &a)

	//指针变量的存储地址
	fmt.Printf("ip变量储存的指针地址: %x\n", ip)

	//使用指针访问值
	fmt.Printf("*ip 变量的值： %d\n", *ip)

	fmt.Println("----------pointer demo2 end----------")

}
