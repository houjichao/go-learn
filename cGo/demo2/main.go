package main

/*
#include <math.h>
*/
import "C"
import "fmt"

/*
在这个Go文件中，我们首先使用import "C"导入C包。我们使用注释引入C头文件math.h，这使得cgo能够找到我们需要的C函数定义。
在main函数中，我们调用了C库中的sqrt函数，将Go变量num转换为C的double类型，然后将C的double类型转换回Go的float64类型，以便将结果输出到控制台。
*/
func main() {
	num := 9.0
	result := C.sqrt(C.double(num))
	fmt.Printf("The square root of %.2f is %.2f\n", num, float64(result))
}
