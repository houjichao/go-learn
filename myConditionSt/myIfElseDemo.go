package myConditionSt

import (
	"fmt"
	"math/rand"
	"time"
)

func IfElseDemo() {
	fmt.Println("----------if else demo start----------")

	//定义局部变量，生成随机数
	rand.Seed(time.Now().Unix())
	var a int = rand.Intn(50)

	/* 判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n")
	} else {
		/* 如果条件为 false 则执行以下语句 */
		fmt.Printf("a 不小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)

	fmt.Println("----------if else demo end----------")

	fmt.Println("----------if ... else if ... else... demo start----------")

	var age int = rand.Intn(100)
	if age == 25 {
		fmt.Println("true")
	} else if age < 25 {
		fmt.Println("too small")
	} else {
		fmt.Println("too big")
	}
	fmt.Println("----------if ... else if ... else... demo end----------")

}
