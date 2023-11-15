package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

func main() {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	fmt.Println(timestamp)
	message := []byte("hello world")
	key := []byte("mysecretkey")

	hash := hmac.New(sha256.New, key)
	hash.Write(message)
	signature := hex.EncodeToString(hash.Sum(nil))

	fmt.Println(signature) // 输出：b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
}
