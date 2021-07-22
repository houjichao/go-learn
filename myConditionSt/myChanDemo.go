package myConditionSt

import "fmt"

/*
chan
引用类型
零值 nil
<发送/写>型 chan<-
<接收/读>型 <-chan
双向型 chan
构造/初始化 make()
关闭 close()
判等 ==
<发送/写>数据 chan <- send_date
<接收/读>数据 recv_data := <- chan
chan 关闭或有数据,读操作不阻塞
chan 未关闭且无数据,读操作阻塞
*/
func ChanDemo() {

	//双向型 chan, 零值 nil
	var ch chan int
	//输入型 chan->
	var ci chan<- int
	//输出型 <-chan
	var co <-chan int
	//make()
	cc := make(chan int, 10)
	ch = cc
	//双向型赋值给单向型正确
	ci = ch
	co = ch
	//单向型赋值给双向型错误
	//ch = ci //
	//ch = co //
	fmt.Println(ci, co, ch) //0xc00008c000 0xc00008c000 0xc00008c000

	//可以赋值 -> 类型兼容 -> 可以判等
	b1 := ch == cc
	b2 := ci == nil
	b3 := ch == ci
	b4 := ch == co
	//不可以赋值 -> 类型不兼容 -> 不可以判等
	// b5 := ci == co //x
	fmt.Println(ci, co, cc, ch) //0xc00008c000 0xc00008c000 0xc00008c000 0xc00008c00
	fmt.Println(b1, b2, b3, b4) //true false true true

	//chan 发送/写
	ch <- 1
	//chan 接收/读
	out := <-ch
	fmt.Println(out) // 1

	ch <- 1
	//close(), chan 关闭
	close(ch) //关闭后读操作不阻塞
	//chan 关闭后,还有数据.读操作,返回 数据 和 true
	out1, ok1 := <-ch
	//chan 关闭后,没有数据.读操作,返回 零值 和 false
	out2, ok2 := <-ch
	fmt.Println(out1, ok1, out2, ok2) // 1 true 0 fals

	ch = make(chan int, 0)
	//chan 没关闭,无数据,读操作阻塞
	//为了程序能跑，加了close关闭
	close(ch)
	out = <-ch

	fmt.Println("上一句阻塞了!!!打印不出这行了?!!!")

}
