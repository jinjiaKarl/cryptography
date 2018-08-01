package main

import (

	"fmt"
	"math/big"
	"crypto/rand"
)

/*
随机种子
伪随机数，是使用一个确定性的算法计算出来的似乎是随机的数序，因此伪随机数实际上并不随机。

那么自然，在计算伪随机数时假如使用的开始值不变的话，那么算法计算出的伪随机数的数序自然也是不变的咯。

这个“开始值”，就被称为随机种子。

可以通过 rand.Seed 方法设置随机种子，如果不设置，则默认值显示为 1，
为了保证每次伪随机数生成器工作时使用的是不同的种子，通常的做法是采用当前时间作为种子。
 */
func main() {
	/*
	// 伪随机数
	rand.Seed(int64(time.Now().Unix()))
	fmt.Println(rand.Intn(100))*/

	//真随机数
	// 如果我们的应用对安全性要求比较高，需要使用真随机数的话，那么可以使用 crypto/rand 包中的方法。
	result, _ := rand.Int(rand.Reader, big.NewInt(100))
	fmt.Println(result)
}
