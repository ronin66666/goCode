## 对称加密

对称加密中，加密用的密钥和解密用的密钥是相同的

常见的对称加密算法有	

- DES

  是一种将64bit明文加密成64bit密文的对称密码算法，密钥长度是56bit，n规格上来说，密钥长度是64bit，但每隔7bit会设置一个用于错误检查的bit，因此密钥长度实质上是56bit

  目前已经可以在短时间内被破解，所以不建议使用

- 3DES

  3DES，将DES重复3次所得到的一种密码算法，也叫做3重DES，处理速度不高，安全性逐渐暴露出问题

  

- AES

  取代DES成为新标准的一种对称密码算法，***ES的密钥长度有128、192、256bit （16、24、32 byte）三种**，目前AES，已经逐步取代DES、3DES，成为首选的对称密码算法

### 对称密码的缺点

不能很好地解决密钥配送问题，在使用对称密码时，需要事先共享密钥

## 公钥密码 - 非对称加密

公钥密码中，密钥分为加密密钥、解密密钥2种，它们并不是同一个密钥

- 加密密钥，一般是公开的，因此该密钥称为公钥（public key）
- 解密密钥，由消息接收者自己保管的，不能公开，因此也称为私钥（private key）
- 公钥和私钥是一 一对应的，是不能单独生成的，一对公钥和密钥统称为密钥对（key pair）
- 由公钥加密的密文，必须使用与该公钥对应的私钥才能解密
- 由私钥加密的密文，必须使用与该私钥对应的公钥才能解密

**解决密钥配送问题**

由消息的接收者，生成一对公钥、私钥，将公钥发给消息的发送者，消息的发送者使用公钥加密消息，消息接收者解密消息

### **RSA**

目前使用最广泛的公钥密码算法是RSA，RSA的名字，由它的3位开发者，即Ron Rivest、Adi Shamir、Leonard Adleman的姓氏首字母组成

**公钥密码的缺点**

加密解密速度比较慢

## 混合密码系统（Hybrid Cryptosystem）

混合密码系统，是将对称密码和公钥密码的优势相结合的方法

**解决了公钥密码速度慢的问题，并通过公钥密码解决了对称密码的密钥配送问题，网络上的密码通信所用的SSL/TLS都运用了混合密码系统**

### 混合密码 - 加密

**会话密钥（session key）**

为本次通信随机生成的临时密钥，作为对称密码的密钥，用于加密消息，提高速度

**加密步骤**（发送消息）

1. 首先，消息发送者要拥有消息接收者的公钥

2. 生成会话密钥，作为对称密码的密钥，加密消息

3. 用消息接收者的公钥，加密会话密钥

4. 将前2步生成的加密结果，一并发给消息接收者

**发送出去的内容包括**

用会话密钥加密的消息（加密方法：对称密码）

用公钥加密的会话密钥（加密方法：公钥密码）

### 混合密码 - 解密

**解密步骤（收到消息）**

1. 消息接收者用自己的私钥解密出会话密钥

2. 再用第1步解密出来的会话密钥，解密消息

### 混合密码 - 加密解密流程

**Alice >>>>> Bob**

- 发送过程，加密过程
  1. Bob先生成一对公钥、私钥
  2. Bob把公钥共享给Alice
  3. Alice随机生成一个会话密钥（临时密钥）
  4. Alice用会话密钥加密需要发送的消息（使用的是对称密码加密）
  5. Alice用Bob的公钥加密会话密钥（使用的是公钥密码加密，也就是非对称密码加密）
  6. Alice把第4、5步的加密结果，一并发送给Bob

- 接收过程，解密过程
  1. Bob利用自己的私钥解密会话密钥（使用的是公钥密码解密，也就是非对称密码解密）
  2. Bob利用会话密钥解密发送过来的消息（使用的是对称密码解密）

## 单向散列函数（One-way **hash** function）

单向散列函数，可以根据根据消息内容计算出散列值，**散列值的长度和消息的长度无关**，无论消息是1bit、10M、100G，单向散列函数都会计算出固定长度的散列值

单向散列函数，又被称为消息摘要函数（message digest function），哈希函数

输出的散列值，也被称为消息摘要（message digest）、指纹（fingerprint）

### 特点

- 根据任意长度的消息，计算出固定长度的散列值
- 计算速度快，能快速计算出散列值
- 消息不同，散列值也不同
- 具备单向性

### 常见的几种单向散列函数

- MD4、MD5

  - 产生128bit的散列值，MD就是Message Digest的缩写，目前已经不安全
  - Mac终端上默认可以使用md5命令

- SHA-1

  产生160bit的散列值，目前已经不安全

- SHA-2

  SHA-256、SHA-384、SHA-512，散列值长度分别是256bit、384bit、512bit

- SHA-3

  全新标准













