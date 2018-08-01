**随机数的应用场景**


下面的场景中就会用到随机数

*   生成密钥

    用于对称密码和消息认证码。

*   生成密钥对

    用于公钥密码和数字签名。

*   生成初始化向量(IV )

    用于分组密码的CBC、CFB和OFB模式。

*   生成nonce

    用于防御重放攻击以及分组密码的CTR模式等。

*   生成盐

    用于基于口令的密码( PBE)等。

**随机数的分类**

随机数主要分为以下三类：

*   弱伪随机数

*   强伪随机数

*   真随机数

随机数的分类是根据随机数的性质进行的分类。随机数的性质分为以下三类：

*   随机性一不存在统计学偏差，是完全杂乱的数列

*   不可预测性一不能从过去的数列推测出”下一个出现的数

*   不可重现性一除非将数列本身保存下来，否则不能重现相同的数列

密码技术中所使用的随机数,仅仅具备随机性是不够的，至少还需要具备不可预测性才行。

![](https://mmbiz.qpic.cn/mmbiz_png/P2euhHCCzXFNibeUeZmHnwiaWmfB0LbnN9FgRrVVNzm9dnrbNnAibCia4WMpWOp0RSOt9PSiaXIDMiafWicviaeq10gkZQ/640?wx_fmt=png&wxfrom=5&wx_lazy=1)