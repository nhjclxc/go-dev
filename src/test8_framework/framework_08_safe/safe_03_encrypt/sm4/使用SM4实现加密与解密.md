在 Golang 中实现 SM4 加密与解密功能，推荐使用国密库 [github.com/tjfoc/gmsm/sm4](https://github.com/tjfoc/gmsm)，它是目前使用最广泛的 SM 系列算法实现之一。

---

## ✅ 一、安装依赖

```bash
go get -u github.com/tjfoc/gmsm/sm4
```

---

## ✅ 二、SM4 加密解密示例（ECB 模式）

```go
package main

import (
	"encoding/hex"
	"fmt"

	"github.com/tjfoc/gmsm/sm4"
)

func main() {
	// 明文
	plaintext := []byte("helloSM4world123") // 必须是 16 字节的倍数
	// 密钥（16字节 = 128位）
	key := []byte("1234567890abcdef")

	// 加密
	ciphertext, err := sm4.Sm4Ecb(key, plaintext, true)
	if err != nil {
		panic(err)
	}
	fmt.Println("密文（hex）:", hex.EncodeToString(ciphertext))

	// 解密
	decrypted, err := sm4.Sm4Ecb(key, ciphertext, false)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密后:", string(decrypted))
}
```

---

## ✅ 三、支持 CBC 模式（带 IV）

```go
package main

import (
	"encoding/hex"
	"fmt"

	"github.com/tjfoc/gmsm/sm4"
)

func main() {
	key := []byte("1234567890abcdef") // 16字节密钥
	iv := []byte("abcdef1234567890") // 16字节IV
	plaintext := []byte("golang-sm4-demo-16bytes") // 必须16字节对齐

	// CBC 加密
	ciphertext, err := sm4.Sm4Cbc(key, iv, plaintext, true)
	if err != nil {
		panic(err)
	}
	fmt.Println("加密后（CBC）:", hex.EncodeToString(ciphertext))

	// CBC 解密
	decrypted, err := sm4.Sm4Cbc(key, iv, ciphertext, false)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密后:", string(decrypted))
}
```

> ⚠️ 注意：SM4 是分组加密算法，加密内容需要 **16 字节对齐（128 位）**，如果不满足，需要做 PKCS#7 填充。

---

## ✅ 四、封装通用工具函数（带自动填充）

如果你需要兼容任意长度数据，可以加上 PKCS#7 填充处理：

```go
import (
	"bytes"
)

func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func pkcs7Unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
```

然后在加密前调用 `pkcs7Padding()`，解密后调用 `pkcs7Unpadding()`。

---

## ✅ 五、补充说明

* **ECB 模式**：简单但不安全，不推荐用于生产。
* **CBC 模式**：常用，但 IV 必须安全管理，不能重复。
* **推荐方案**：使用 CBC + PKCS#7 填充。
* **密钥长度**：固定为 16 字节（128 位）。

---

如你想要用在 Gin 接口前后端加解密，可以继续说“前端 JS 怎么配合 SM4 加密，如何嵌入 Gin 中间件”等，我可以帮你继续补充。
