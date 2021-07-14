package myConst

import "fmt"

/*
iota，特殊常量，可以认为是一个可以被编译器修改的常量。
iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
iota 可以被用作枚举值：
const (
    a = iota
    b = iota
    c = iota
)
第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：
const (
    a = iota
    b
    c
)
*/
/*
iota 只是在同一个 const 常量组内递增，每当有新的 const 关键字时，iota 计数会重新开始。
const (
    i = iota
    j = iota
    x = iota
)
const xx = iota
const yy = iota
func main(){
    println(i, j, x, xx, yy)
}

// 输出是 0 1 2 0 0
*/
func ConstTest2()  {
	const (
		a = iota  //0
		b         //1
		c         //2
		d = "ha"  //ha
		e         //ha
		f = 100   //100
		g         //100
		h = iota  //7,恢复计数
		i         //8
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)
}
