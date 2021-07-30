package myLanguageConversion

import (
	"fmt"
	"strconv"
)

/**
语言类型转换
类型转换用于将一种数据类型的变量转换为另外一种类型的变量。Go 语言类型转换基本格式如下：
type_name(expression)
type_name 为类型，expression 为表达式。

 那么float32和float64有什么区别呢?
数位的区别
一个在内存中占分别32和64个bits，也就是4bytes或8bytes
数位越高浮点数的精度越高
它会影响深度学习计算效率?
float64占用的内存是float32的两倍，是float16的4倍；比如对于CIFAR10数据集，如果采用float64来表示，
需要60000*32*32*3*8/1024**3=1.4G，光把数据集调入内存就需要1.4G；如果采用float32，只需要0.7G，如果采用float16，只需要0.35G左右；
占用内存的多少，会对系统运行效率有严重影响；（因此数据集文件都是采用uint8来存在数据，保持文件最小）

*/
func LanguageConversionDemo() {

	fmt.Println("----------language conversion demo start----------")

	var sum int = 17
	var count int = 5
	var mean float64

	mean = float64(sum) / float64(count)
	fmt.Printf("mean的值为%f\n", mean)

	/*
		strconv.FormatFloat(f,fmt,prec,bitSize)
		f：要转换的浮点数
		fmt：格式标记（b、e、E、,f、g、G）
		prec：精度（数字部分的长度，不包括指数部分）
		bitSize：指定浮点类型（32:float32、64:float64）

		‘b’ (-ddddp±ddd，二进制指数)
		‘e’ (-d.dddde±dd，十进制指数)
		‘E’ (-d.ddddE±dd，十进制指数)
		‘f’ (-ddd.dddd，没有指数)
		‘g’ (‘e’:大指数，’f’:其它情况)
		‘G’ (‘E’:大指数，’f’:其它情况)
		如果格式标记为 ‘e’，’E’和’f’，则 prec 表示小数点后的数字位数
		如果格式标记为 ‘g’，’G’，则 prec 表示总的数字位数（整数部分+小数部分）
	*/
	//返回的string类型，不能用%f格式化
	newMean := strconv.FormatFloat(mean, 'f', 1, 64)
	fmt.Printf("mean的值为%s\n", newMean)

	/*
		go 不支持隐式转换类型
		会报错：cannot use a (type int64) as type int32 in assignment
	*/
	var a int64 = 3
	var b int32
	//b = a
	b = int32(a)
	fmt.Printf("b 为 ： %d\n", b)

	fmt.Println("----------language conversion demo end----------")

}
