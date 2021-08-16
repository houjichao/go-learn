package myPointer

import "fmt"

/**
空指针
*/
func PointerDemo3() {

	fmt.Println("----------pointer demo3 start----------")
	var ptr *int
	fmt.Printf("ptf的值为：%x\n", ptr)
	if ptr == nil {
		fmt.Println("ptr is null")
	} else {
		fmt.Println("ptr is not null")
	}
	fmt.Println("----------pointer demo3 end----------")

}
