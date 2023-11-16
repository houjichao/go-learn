package main

/*
#include "sum.h"
*/
import "C"
import "fmt"

func main() {
	a := 10
	b := 20
	sum := C.add1(C.int(a), C.int(b))
	fmt.Printf("The sum of %d and %d is %d\n", a, b, int(sum))
}
