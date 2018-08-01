package main

import (
	"fmt"
	"golang.org/x/crypto/ripemd160"
)

// 160bit->40位16进制字符。
func main() {
	hasher := ripemd160.New()
	hasher.Write([]byte("The quick brown fox jumps over the lazy dog"))
	hashBytes := hasher.Sum(nil)
	hashString := fmt.Sprintf("%x", hashBytes)
	fmt.Println(hashString)
}