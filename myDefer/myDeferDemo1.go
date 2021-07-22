package myDefer

import "os"

/*
处理业务或逻辑中涉及成对的操作是一件比较烦琐的事情，比如打开和关闭文件、接收请求和回复请求、加锁和解锁等。在这些操作中，最容易忽略的就是在每个函数退出处正确地释放和关闭资源。
defer 语句正好是在函数退出时执行的语句，所以使用 defer 能非常方便地处理资源释放问题。
*/
func FileSize(fileName string) int64 {
	// 根据文件名打开文件, 返回文件句柄和错误
	f, err := os.Open(fileName)
	// 如果打开时发生错误, 返回文件大小为0
	if err != nil {
		return 0
	}
	// 取文件状态信息
	info, err := f.Stat()

	// 如果获取信息时发生错误, 关闭文件并返回文件大小为0
	if err != nil {
		f.Close()
		return 0
	}
	// 取文件大小
	size := info.Size()
	// 关闭文件
	f.Close()

	// 返回文件大小
	return size
}

func DeferFileSize(fileName string) int64 {
	f, err := os.Open(fileName)
	if err != nil {
		return 0
	}
	//延迟调用close,此时close不会被调用
	defer f.Close()

	/*
		在文件正常打开后，使用 defer，将 f.Close() 延迟调用，注意，不能将这一句代码放在os.Open后，
		一旦文件打开错误，f 将为空，在延迟语句触发时，将触发宕机错误。
	*/

	// 取文件状态信息
	info, err := f.Stat()
	if err != nil {
		//defer机制自动触发，调用close关闭文件
		return 0
	}

	size := info.Size()

	//defer机制自动触发，调用close关闭文件
	return size

}
