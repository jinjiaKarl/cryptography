常用的密码学：

*   散列函数(也称哈希函数)算法：MD5、SHA

*   对称加密算法：DES、3DES、AES

*   非对称加密算法：DSA、RSA、ECC


## 密码学参考资料

[密码学介绍](http://blog.51cto.com/11821908/category1.html)

[非对称加密技术- RSA算法数学原理分析](https://learnblockchain.cn/2017/11/15/asy-encryption/)

[椭圆曲线加密算法和国密算法](https://zhuanlan.zhihu.com/p/36326221)

[golanger/Golang学习笔记](https://www.teakki.com/p/57df64fbda84a0c4533817a6)

[Go 密码学应用 - George Tankersley](https://blog.lab99.org/post/golang-2017-09-23-video-go-for-crypto-developers.html)

[MD5算法原理与实现](https://blog.csdn.net/xiaofengcanyuexj/article/details/37698801)

[hash算法原理详解](https://www.jianshu.com/p/f9239c9377c5)

[Hash算法总结](https://www.jianshu.com/p/bf1d7eee28d0)

[哈希算法(go版)](https://my.oschina.net/ifraincoat/blog/604415)

[图说区块链](https://xiaoxueying.gitbooks.io/graphic-cryptology/content/block_cipher_mode.html)

[golang中实现RSA(PKCS#1)加密解密](https://blog.csdn.net/yue7603835/article/details/73433617)



# 什么是 hash 算法

　　散列方法的主要思想是根据结点的关键码值来确定其存储地址：以关键码值K为自变量，通过一定的函数关系h(K)(称为散列函数)，计算出对应的函数值来，把这个值解释为结点的存储地址，将结点存入到此存储单元中。检索时，用同样的方法计算地址，然后到相应的单元里去取要找的结点。通过散列方法可以对结点进行快速检索。散列(hash，也称“哈希”)是一种重要的存储方式，也是一种常见的检索方法。

　　散列算法(Hash Algorithm)，又称哈希算法，杂凑算法，是一种从任意文件中创造小的数字「指纹」的方法。与指纹一样，散列算法就是一种以较短的信息来保证文件唯一性的标志，这种标志与文件的每一个字节都相关，而且难以找到逆向规律。因此，当原有文件发生改变时，其标志值也会发生改变，从而告诉文件使用者当前的文件已经不是你所需求的文件。

　　若结构中存在和关键字K相等的记录，则必定在f(K)的存储位置上。由此，不需比较便可直接取得所查记录。称这个对应关系f为散列函数(Hash function)，按这个事先建立的表为散列表。

　　对不同的关键字可能得到同一散列地址，即key1≠key2，而f(key1)=f(key2)，这种现象称碰撞。具有相同函数值的关键字对该散列函数来说称做同义词。综上所述，根据散列函数H(key)和处理冲突的方法将一组关键字映象到一个有限的连续的地址集(区间)上，并以关键字在地址集中的“象” 作为记录在表中的存储位置，这种表便称为散列表，这一映象过程称为散列造表或散列，所得的存储位置称散列地址。

　　若对于关键字集合中的任一个关键字，经散列函数映象到地址集合中任何一个地址的概率是相等的，则称此类散列函数为均匀散列函数(Uniform Hash function)，这就是使关键字经过散列函数得到一个“随机的地址”，从而减少冲突。


# Hash算法有什么特点


　　一个优秀的 hash 算法，将能实现：

　　正向快速：给定明文和 hash 算法，在有限时间和有限资源内能计算出 hash 值。

　　逆向困难：给定(若干) hash 值，在有限时间内很难(基本不可能)逆推出明文。

　　输入敏感：原始输入信息修改一点信息，产生的 hash 值看起来应该都有很大不同。

　　冲突避免：很难找到两段内容不同的明文，使得它们的 hash 值一致(发生冲突)。即对于任意两个不同的数据块，其hash值相同的可能性极小;对于一个给定的数据块，找到和它hash值相同的数据块极为困难。

　　但在不同的使用场景中，如数据结构和安全领域里，其中对某一些特点会有所侧重。


# 常用Hash函数


直接取余法：f(x):= x mod maxM ; maxM一般是不太接近 2^t 的一个质数。

乘法取整法：f(x):=trunc((x/maxX)*maxlongit) mod maxM，主要用于实数。

平方取中法：f(x):=(x*x div 1000 ) mod 1000000); 平方后取中间的，每位包含信息比较多。


# Hash构造方法


　　散列函数能使对一个数据序列的访问过程更加迅速有效，通过散列函数，数据元素将被更快地定位。

　　(详细构造方法可以参考hash函数中的【哈希表的构造方法】)

　　1.直接寻址法：取关键字或关键字的某个线性函数值为散列地址。即H(key)=key或H(key) = a·key + b，其中a和b为常数(这种散列函数叫做自身函数)

　　2\. 数字分析法

　　3\. 平方取中法

　　4\. 折叠法

　　5\. 随机数法

　　6\. 除留余数法：取关键字被某个不大于散列表表长m的数p除后所得的余数为散列地址。即 H(key) = key MOD p,p<=m。不仅可以对关键字直接取模，也可在折叠、平方取中等运算之后取模。对p的选择很重要，一般取素数或m，若p选的不好，容易产生同义词。



# Hash算法是如何实现的


　　作为散列算法，首要的功能就是要使用一种算法把原有的体积很大的文件信息用若干个字符来记录，还要保证每一个字节都会对最终结果产生影响。那么大家也许已经想到了，求模这种算法就能满足我们的需要。

　　事实上，求模算法作为一种不可逆的计算方法，已经成为了整个现代密码学的根基。只要是涉及到计算机安全和加密的领域，都会有模计算的身影。散列算法也并不例外，一种最原始的散列算法就是单纯地选择一个数进行模运算，比如以下程序。

```
#  构造散列函数
def hash(a):
    return a % 8

#  测试散列函数功能
print(hash(233))
print(hash(234))
print(hash(235))

# 输出结果
- 1
- 2
```

很显然，上述的程序完成了一个散列算法所应当实现的初级目标：用较少的文本量代表很长的内容（求模之后的数字肯定小于8）。但也许你已经注意到了，单纯使用求模算法计算之后的结果带有明显的规律性，这种规律将导致算法将能难保证不可逆性。所以我们将使用另外一种手段，那就是异或。

再来看下面一段程序，我们在散列函数中加入一个异或过程。

```
#  构造散列函数
def hash(a):
    return (a % 8) ^ 5

#  测试散列函数功能
print(hash(233))
print(hash(234))
print(hash(235))

# 输出结果
- 4
- 7
- 6
```


      很明显的，加入一层异或过程之后，计算之后的结果规律性就不是那么明显了。

      当然，大家也许会觉得这样的算法依旧很不安全，如果用户使用连续变化的一系列文本与计算结果相比对，就很有可能找到算法所包含的规律。但是我们还有其他的办法。比如在进行计算之前对原始文本进行修改，或是加入额外的运算过程（如移位），比如以下程序。

```
#  构造散列函数
def hash(a):
    return (a + 2 + (a << 1)) % 8 ^ 5

#  测试散列函数功能
print(hash(233))
print(hash(234))
print(hash(235))

# 输出结果
- 0
- 5
- 6
```

      这样处理得到的散列算法就很难发现其内部规律，也就是说，我们并不能很轻易地给出一个数，让它经过上述散列函数运算之后的结果等于4——除非我们去穷举测试。
 

# 处理冲突的方法


　　1.开放寻址法;Hi=(H(key) + di) MOD m,i=1,2,…，k(k<=m-1)，其中H(key)为散列函数，m为散列表长，di为增量序列，可有下列三种取法：

　　1). di=1,2,3,…，m-1，称线性探测再散列;

　　2). di=1^2,(-1)^2,2^2,(-2)^2,(3)^2,…，±(k)^2,(k<=m/2)称二次探测再散列;

　　3). di=伪随机数序列，称伪随机探测再散列。

　　2\. 再散列法：Hi=RHi(key),i=1,2,…，k RHi均是不同的散列函数，即在同义词产生地址冲突时计算另一个散列函数地址，直到冲突不再发生，这种方法不易产生“聚集”，但增加了计算时间。

　　3\. 链地址法(拉链法)

　　4\. 建立一个公共溢出区



# Hash有哪些流行的算法


  目前流行的 Hash 算法包括 MD5、SHA-1 和 SHA-2。

*   MD4（RFC 1320）是 MIT 的 Ronald L. Rivest 在 1990 年设计的，MD 是 Message Digest 的缩写。其输出为 128 位。MD4 已证明不够安全。

*   MD5（RFC 1321）是 Rivest 于1991年对 MD4 的改进版本。它对输入仍以 512 位分组，其输出是 128 位。MD5 比 MD4 复杂，并且计算速度要慢一点，更安全一些。MD5 已被证明不具备"强抗碰撞性"。

*   SHA （Secure Hash Algorithm）是一个 Hash 函数族，由 NIST（National Institute of Standards and Technology）于 1993 年发布第一个算法。目前知名的 SHA-1 在 1995 年面世，它的输出为长度 160 位的 hash 值，因此抗穷举性更好。SHA-1 设计时基于和 MD4 相同原理，并且模仿了该算法。SHA-1 已被证明不具"强抗碰撞性"。

   为了提高安全性，NIST 还设计出了 SHA-224、SHA-256、SHA-384，和 SHA-512 算法（统称为 SHA-2），跟 SHA-1 算法原理类似。SHA-3 相关算法也已被提出。


# Hash算法的「碰撞」


　　在实现算法章节的第一个例子，我们尝试的散列算法得到的值一定是一个不大于8的自然数，因此，如果我们随便拿9个数去计算，肯定至少会得到两个相同的值，我们把这种情况就叫做散列算法的「碰撞」(Collision)。

　　这很容易理解，因为作为一种可用的散列算法，其位数一定是有限的，也就是说它能记录的文件是有限的——而文件数量是无限的，两个文件指纹发生碰撞的概率永远不会是零。

　　但这并不意味着散列算法就不能用了，因为凡事都要考虑代价，买光所有彩票去中一次头奖是毫无意义的。现代散列算法所存在的理由就是，它的不可逆性能在较大概率上得到实现，也就是说，发现碰撞的概率很小，这种碰撞能被利用的概率更小。

　　随意找到一组碰撞是有可能的，只要穷举就可以。散列算法得到的指纹位数是有限的，比如MD5算法指纹字长为128位，意味着只要我们穷举21282128次，就肯定能得到一组碰撞——当然，这个时间代价是难以想象的，而更重要的是，仅仅找到一组碰撞并没有什么实际意义。更有意义的是，如果我们已经有了一组指纹，能否找到一个原始文件，让它的散列计算结果等于这组指纹。如果这一点被实现，我们就可以很容易地篡改和伪造网络证书、密码等关键信息。

　　你也许已经听过MD5已经被破解的新闻——但事实上，即便是MD5这种已经过时的散列算法，也很难实现逆向运算。我们现在更多的还是依赖于海量字典来进行尝试，也就是通过已经知道的大量的文件——指纹对应关系，搜索某个指纹所对应的文件是否在数据库里存在。


# 哈希函数



　　(1)余数法:先估计整个哈希表中的表项目数目大小。然后用这个估计值作为除数去除每个原始值，得到商和余数。用余数作为哈希值。因为这种方法产生冲突的可能性相当大，因此任何搜索算法都应该能够判断冲突是否发生并提出取代算法。

　　(2)折叠法:这种方法是针对原始值为数字时使用，将原始值分为若干部分，然后将各部分叠加，得到的最后四个数字(或者取其他位数的数字都可以)来作为哈希值。

　　(3)基数转换法:当原始值是数字时，可以将原始值的数制基数转为一个不同的数字。例如，可以将十进制的原始值转为十六进制的哈希值。为了使哈希值的长度相同，可以省略高位数字。

　　(4)数据重排法:这种方法只是简单的将原始值中的数据打乱排序。比如可以将第三位到第六位的数字逆序排列，然后利用重排后的数字作为哈希值。

　　哈希函数并不通用，比如在数据库中用能够获得很好效果的哈希函数，用在密码学或错误校验方面就未必可行。在密码学领域有几个著名的哈希函数。这些函数包括MD2、MD4以及MD5，利用散列法将数字签名转换成的哈希值称为信息摘要(message-digest)，另外还有安全散列算法(SHA)，这是一种标准算法，能够生成更大的(60bit)的信息摘要，有点儿类似于MD4算法。

