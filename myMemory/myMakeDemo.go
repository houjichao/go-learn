package myMemory

import "fmt"

/*
golang 分配内存主要有内置函数new和make，今天我们来探究一下make有哪些玩法。

map只能为slice, map, channel(管道)分配内存，并返回一个初始化的值。首先来看下make有以下三种不同的用法：
slice:
slice表示切片(分片)，例如对一个数组进行切片，取出数组中的一部分值。在现代编程语言中，slice(切片)几乎成为一种必备特性，
它可以从一个数组(列表)中取出任意长度的子数组(列表)，为操作数据结构带来非常大的便利性
Go中的slice不仅仅是一种切片动作，还是一种数据结构(就像数组一样)。
Go中的slice依赖于数组，它的底层就是数组，所以数组具有的优点，slice都有。且slice支持可以通过append向slice中追加元素，长度不够时会动态扩展，通过再次slice切片，可以得到得到更小的slice结构，可以迭代、遍历等。
实际上slice是这样的结构：先创建一个有特定长度和数据类型的底层数组，然后从这个底层数组中选取一部分元素，返回这些元素组成的集合(或容器)，
并将slice指向集合中的第一个元素。换句话说，slice自身维护了一个指针属性，指向它底层数组中的某些元素的集合。
例如，初始化一个slice数据结构：
my_slice := make([]int, 3，5)
// 输出slice
fmt.Println(my_slice)    // 输出：[0 0 0]

1. make(map[string]string)

2. make([]int, 2)

3. make([]int, 2, 4)

1. 第一种用法，即缺少长度的参数，只传类型，这种用法只能用在类型为map或chan的场景，例如make([]int)是会报错的。这样返回的空间长度都是默认为0的。

2. 第二种用法，指定了长度，例如make([]int, 2)返回的是一个长度为2的slice

3. 第三种用法，第二参数指定的是切片的长度，第三个参数是用来指定预留的空间长度，例如a := make([]int, 2, 4),
这里值得注意的是返回的切片a的总长度是4，预留的意思并不是另外多出来4的长度，其实是包含了前面2个已经切片的个数的。所以举个例子当你这样用的时候 a := make([]int, 4, 2)，就会报语法错误。

因此，当我们为slice分配内存的时候，应当尽量预估到slice可能的最大长度，
通过给make传第三个参数的方式来给slice预留好内存空间，这样可以避免二次分配内存带来的开销，大大提高程序的性能。

而事实上，我们其实是很难预估切片的可能的最大长度的，这种情况下，当我们调用append为slice追加元素时，
golang为了尽可能的减少二次分配内存，并不是每一次都只增加一个单位的内存空间，而且遵循这样一种扩容机制：

当有预留的未使用的空间时，直接对未使用的空间进行切片追加，当预留的空间全部使用完毕的时候，扩容的空间将会是当前的slice长度的一倍，
例如当前slice的长度为4，进行一次append操作之后，cap(a)返回的长度将会是8.来看下面这段演示代码:
*/
func MakeDemo() {
	fmt.Println("----------make demo start----------")

	a := make([]int, 0)
	n := 20
	for i := 0; i < n; i++ {
		a = append(a, 1)
		fmt.Printf("len=%d cap=%d \n", len(a), cap(a))
	}
	fmt.Println("----------make demo end----------")

	/*
		Output:
		len=1 cap=1  // 第一次扩容
		len=2 cap=2 // 第二次扩容
		len=3 cap=4 // 第三次扩容
		len=4 cap=4
		len=5 cap=8 // 第四次扩容
		len=6 cap=8
		len=7 cap=8
		len=8 cap=8
		len=9 cap=16 // 第五次扩容
		len=10 cap=16
		len=11 cap=16
		len=12 cap=16
		len=13 cap=16
		len=14 cap=16
		len=15 cap=16
		len=16 cap=16
		len=17 cap=32 // 第六次扩容
		len=18 cap=32
		len=19 cap=32
		len=20 cap=32
	*/

}
