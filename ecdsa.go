package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"crypto/rand"
	"math/big"
	"fmt"
)

/*
ECDSA的全名是Elliptic Curve DSA，即椭圆曲线DSA。
它是Digital Signature Algorithm (DSA)应用了椭圆曲线加密算法的变种。
椭圆曲线算法的原理很复杂，但是具有很好的公开密钥算法特性，通过公钥无法逆向获得私钥。

签名过程

假设要签名的消息是一个字符串：“Hello World!”。DSA签名的第一个步骤是对待签名的消息生成一个消息摘要。
不同的签名算法使用不同的消息摘要算法。而ECDSA256使用SHA256生成256比特的摘要。
摘要生成结束后，应用签名算法对摘要进行签名： 产生一个随机数k 利用随机数k，计算出两个大数r和s。将r和s拼在一起就构成了对消息摘要的签名。
这里需要注意的是，因为随机数k的存在，对于同一条消息，使用同一个算法，产生的签名是不一样的。
从函数的角度来理解，签名函数对同样的输入会产生不同的输出。因为函数内部会将随机值混入签名的过程。

验证过程 关于验证过程，这里不讨论它的算法细节。
从宏观上看，消息的接收方从签名中分离出r和s，然后利用公开的密钥信息和s计算出r。
如果计算出的r和接收到的r值相同，则表示验证成功。否则，表示验证失败。
 */


func main() {
	// 明文
	message := []byte("Hello world")

	key, err := NewSignningKey()
	if err != nil {
		return
	}

	signature, err := Sign(message, key)

	fmt.Printf("签名后：%x\n", signature)
	if err != nil {
		return
	}

	if !Verify(message, signature, &key.PublicKey) {
		fmt.Println("验证失败！")
		return
	}else{
		fmt.Println("验证成功！")
	}
}

func NewSignningKey() (*ecdsa.PrivateKey, error)  {
	key, err := ecdsa.GenerateKey(elliptic.P256(),rand.Reader)
	return  key, err
}

// Sign signs arbitrary data using ECDSA.
func Sign(data []byte, privkey *ecdsa.PrivateKey) ([]byte, error) {
	// hash message
	digest := sha256.Sum256(data)

	// sign the hash
	r, s, err := ecdsa.Sign(rand.Reader, privkey, digest[:])
	if err != nil {
		return nil, err
	}

	// encode the signature {R, S}
	// big.Int.Bytes() will need padding in the case of leading zero bytes
	params := privkey.Curve.Params()
	curveOrderByteSize := params.P.BitLen() / 8
	rBytes, sBytes := r.Bytes(), s.Bytes()
	signature := make([]byte, curveOrderByteSize*2)
	copy(signature[curveOrderByteSize-len(rBytes):], rBytes)
	copy(signature[curveOrderByteSize*2-len(sBytes):], sBytes)

	return signature, nil
}

// Verify checks a raw ECDSA signature.
// Returns true if it's valid and false if not.
func Verify(data, signature []byte, pubkey *ecdsa.PublicKey) bool {
	// hash message
	digest := sha256.Sum256(data)

	curveOrderByteSize := pubkey.Curve.Params().P.BitLen() / 8

	r, s := new(big.Int), new(big.Int)
	r.SetBytes(signature[:curveOrderByteSize])
	s.SetBytes(signature[curveOrderByteSize:])

	return ecdsa.Verify(pubkey, digest[:], r, s)
}