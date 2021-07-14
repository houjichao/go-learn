package main

import (
	"fmt"
	"github.com/hjc/go-learn/myConst"
	mathClass "github.com/hjc/go-learn/myMath"
	"github.com/hjc/go-learn/myStr"
	"github.com/hjc/go-learn/myVar"
	"strconv"
)

var isActive bool  // 全局变量声明
var enabled, disabled = true, false  // 忽略类型的声明

func main() {
	/*
		当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，
		那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；
		标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。
	*/
	fmt.Println("Hello World")
	//一行代表一个语句结束，不用分号结尾
	//如果你打算将多个语句写在同一行，它们则必须使用 ; 人为区分，但在实际开发中我们并不鼓励这种做法。
	fmt.Println("Hello Go")
	//字符串链接使用加号
	fmt.Println("Google" + "Search")

	//变量声明必须使用空格隔开
	var age int
	age = 18
	fmt.Println("我的年龄是" + strconv.Itoa(age))

	//格式化字符串
	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)

	/**
	关于包，根据本地测试得出以下几点：
	文件名与包名没有直接关系，不一定要将文件名与包名定成同一个。
	文件夹名与包名没有直接关系，并非需要一致。
	同一个文件夹下的文件只能有一个包名，否则编译报错。
	*/
	fmt.Println(mathClass.Add(1, 1))
	fmt.Println(mathClass.Sub(1, 1))

	test()

	myStr.Test()

	//变量
	myVar.VarTest()
	myVar.VarTest2()
	myVar.VarTest3()

	//常量
	myConst.ConstTest1()
	myConst.ConstTest2()
	myConst.ConstTest3()

}

func test() {
	var available bool  // 一般声明
	valid := false      // 简短声明
	available = true    // 赋值操作

	fmt.Println(available)
	fmt.Println(valid)
}


