package myConditionSt

import (
	"fmt"
	"math/rand"
	"time"
)

func SwitchDemo1() {

	fmt.Println("----------switch demo1 start----------")


	/* 定义局部变量 */
	var grade string = "B"
	rand.Seed(time.Now().Unix())
	var marks int = rand.Intn(100)

	switch  {
	case marks > 90:
		grade = "A"
	case marks > 80:
		grade = "B"
	case marks > 60:
		grade = "C"
	default:
		grade = "D"

	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)

	fmt.Println("----------switch demo1 end----------")


}
