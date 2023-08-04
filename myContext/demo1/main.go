package main

import (
	"context"
	"fmt"
	"time"
)

/**
我们可以通过一个代码片段了解 context.Context 是如何对信号进行同步的。
在这段代码中，我们创建了一个过期时间为 1s 的上下文，并向上下文传入 handle 函数，该方法会使用 500ms 的时间处理传入的请求：
因为过期时间大于处理时间，所以我们有足够的时间处理该请求，运行上述代码会打印出下面的内容：
$ go run context.go
process request with 500ms
main context deadline exceeded
handle 函数没有进入超时的 select 分支，但是 main 函数的 select 却会等待 context.Context
超时并打印出 main context deadline exceeded。

如果我们将处理请求时间增加至 1500ms，整个程序都会因为上下文的过期而被中止，：
$ go run context.go
main context deadline exceeded
handle context deadline exceeded
Go
相信这两个例子能够帮助各位读者理解 context.Context 的使用方法和设计原理 —
多个 Goroutine 同时订阅 ctx.Done() 管道中的消息，一旦接收到取消信号就立刻停止当前正在执行的工作。

*/

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
