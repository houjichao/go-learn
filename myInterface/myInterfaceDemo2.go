package myInterface

import "fmt"

/**
给接口增加参数
*/

type Person interface {
	name() string
	age() int
}

type Man struct {
}

func (man Man) name() string {
	return "houjichao"
}

func (man Man) age() int {
	return 27
}

type Woman struct {
}

func (woman Woman) name() string {
	return "gaoxin"
}

func (woman Woman) age() int {
	return 25
}

func InterfaceDemo2() {
	fmt.Println("----------interface demo2 start----------")

	var person Person

	person = new(Woman)
	fmt.Println(person.name())
	fmt.Println(person.age())

	person = new(Man)
	fmt.Println(person.name())
	fmt.Println(person.age())

	fmt.Println("----------interface demo2 end----------")

}
