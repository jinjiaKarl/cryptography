package main

import (
	"crypto/dsa"
	"crypto/rand"
	"fmt"
)


//数字签名算法（DSA，Digital Signature Algorithm），
// 是一种公开密钥算法，不能用于加密，只能用于数字签名。
// 主要用作为接收者验证数字的完整性和数据发送者的身份，
// DSA算法的安全性基于解离散对数的困难性。
func main() {
	var params dsa.Parameters

	//生成参数
	if e := dsa.GenerateParameters(&params, rand.Reader, dsa.L1024N160); e != nil {
		fmt.Println(e)
	}

	//生成私钥
	var priKey dsa.PrivateKey

	priKey.Parameters = params
	if e := dsa.GenerateKey(&priKey, rand.Reader); e != nil {
		fmt.Println(e)
	}

	// 根据私钥生成公钥
	pubKey := priKey.PublicKey

	// 消息
	message := []byte("hello world")

	// 使用私钥进行签名，产生整数对(r,s)
	r, s, e := dsa.Sign(rand.Reader,&priKey, message)
	if e != nil {
		fmt.Println(e)
	}

	// 认证
	fmt.Printf("认证 %q (r:%s, s:%s)\n", message, r, s)
	if dsa.Verify(&pubKey,message,r,s) {
		fmt.Println("认证成功!")
	}else {
		fmt.Println("认证失败!")
	}


}
