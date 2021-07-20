package myConditionSt

import (
	"fmt"
	"math/rand"
	"time"
)

/**
if嵌套
*/
func IfNestDemo() {
	fmt.Println("----------if nest demo start----------")

	rand.Seed(time.Now().Unix())
	var a int = rand.Intn(500)
	var b int = rand.Intn(500)

	if a == 100 {
		if b == 200 {
			fmt.Printf("a的值为100，b的值为200\n")
		}
	}
	fmt.Printf("a 值为 : %d\n", a)
	fmt.Printf("b 值为 : %d\n", b)
	fmt.Println("----------if nest demo end----------")

	fmt.Println("----------判断用户密码输入 demo start----------")

	fmt.Printf("请输入密码：   \n")
	//fmt.Scan(&a)
	if a == 5211314 {
		fmt.Println("请再次输入密码：")
		fmt.Scan(&b)
		if b == 5211314 {
			fmt.Println("密码正确，门锁已打开")
		} else {
			fmt.Println("非法入侵，已自动报警")
		}
	} else {
		fmt.Println("非法入侵，已自动报警")
	}

	fmt.Println("----------判断用户密码输入 demo end----------")


}
