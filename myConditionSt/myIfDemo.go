package myConditionSt

import (
	"fmt"
	"math/rand"
	"time"
)

/*
if语句
*/
func IfDemo() {

	fmt.Println("----------if demo start----------")
	//定义局部变量
	/*
		在本环境中，不论编译运行多少次，都是这个输出。

		为什么没有产生随机的效果呢？

		此程序的运行环境是固定的，因此 rand.Intn 总是会返回相同的数字。 （要得到不同的数字，需为生成器提供不同的种子数）
	*/
	//rand.Seed(time.Now().UnixNano())作用
	//获取随机数，不加随机种子，每次遍历获取都是重复的一些随机数据
	//设置随机数种子，加上这行代码，可以保证每次随机都是随机的
	rand.Seed(time.Now().Unix())
	//如果不加随机数种子，那么每次的随机数都是一样的
	//fmt.Println("My first lucky number is", rand.Intn(10))
	//fmt.Println("My senond lucky number is", rand.Intn(10))
	var a int = rand.Intn(50)
	// 使用if语句判断布尔表达式
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)

	fmt.Println("----------if demo end----------")

}
