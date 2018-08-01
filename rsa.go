package main

import (
	"crypto/rsa"
	"crypto/rand"
	"fmt"
	"crypto/md5"
	"crypto"
)

/*
公钥加密算法于1987年首次公开，RSA是提出这个算法的三人姓氏开头字母组成，可用于加密，也可以用于数字签名。RSA的安全性基于大数分解的困难性。

加密算法：

最优非对称加密填充（OAEP，Optimal Asymmetric Encryption Padding），在随机预言模型下，用来处理非对称加密前的明文；

公钥密码学标准（PKCS，The Public-Key Cryptography Standards），是由美国RSA数据安全公司及其合作伙伴制定的一组公钥密码学标准，其中包括证书申请、证书更新、证书作废表发布、扩展证书内容以及数字签名、数字信封的格式等方面的一系列相关协议。

签名认证：

公钥密码学标准（PKCS）；

概率签名方案（PSS，Probabilistic Signature Scheme），与PKCS不同的是，它支持添加盐（Salt）。
 */
func main() {
	//生成私钥
	priKey, e := rsa.GenerateKey(rand.Reader,1024)
	if e != nil {
		fmt.Println(e)
	}

	//根据私钥产生公钥
	pubKey := &priKey.PublicKey

	//明文
	plaintext := []byte("hello world")

	//加密生成密文
	fmt.Printf("%q\n加密:\n", plaintext)
	ciphertext, e := rsa.EncryptOAEP(md5.New(), rand.Reader, pubKey, plaintext, nil)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("\t%x\n", ciphertext)

	//解密得到明文
	fmt.Printf("解密:\n")
	plaintext, e = rsa.DecryptOAEP(md5.New(), rand.Reader, priKey, ciphertext, nil)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("\t%q\n", plaintext)

	//消息先进行hash处理
	h := md5.New()
	h.Write(plaintext)
	hashed :=h.Sum(nil)
	fmt.Printf("%q MD5 Hashed:\n\t%x\n", plaintext, hashed)

	//签名
	opts := &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto, Hash: crypto.MD5}
	sig, e := rsa.SignPSS(rand.Reader, priKey, crypto.MD5, hashed, opts)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("签名:\n\t%x\n", sig)

	//认证
	fmt.Printf("验证结果:")
	if e := rsa.VerifyPSS(pubKey, crypto.MD5, hashed, sig, opts); e != nil {
		fmt.Println("失败:", e)
	} else {
		fmt.Println("成功.")
	}

}
