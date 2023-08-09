package demo1

import (
	"fmt"
	"testing"
)

func Test_Chunk(t *testing.T) {
	var chunks [][]string
	for i, chunk := range chunks {
		fmt.Printf("Chunk %d: %v\n", i+1, chunk)
	}

	strs := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	// 将字符串数组均分为5个一组
	chunkSize := 5
	chunks = Chunk(strs, chunkSize)
	// 打印结果
	for i, chunk := range chunks {
		fmt.Printf("Chunk %d: %v\n", i+1, chunk)
	}
}
